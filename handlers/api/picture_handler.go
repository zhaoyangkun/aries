package api

import (
	"aries/forms"
	"aries/log"
	"aries/models"
	"aries/utils"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"

	"strconv"
	"time"
)

const (
	qubuUploadURL  = "https://7bu.top/api/v1/upload"
	smmsUploadURL  = "https://sm.ms/api/v2/upload"
	imgbbUploadURL = "https://api.imgbb.com/1/upload"
)

var (
	errFileUpload      = errors.New("文件上传失败")
	errCosClientCreate = errors.New("存储桶创建失败")
)

type PictureHandler struct {
}

// GetPicturesByPage
// @Summary 分页获取图片
// @Tags 图床
// @version 1.0
// @Accept application/json
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Param key query string false "关键词"
// @Param storage_name query string false "存储类型"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/images [get]
func (p *PictureHandler) GetPicturesByPage(ctx *gin.Context) {
	pageForm := forms.PicturePageForm{}
	_ = ctx.ShouldBindQuery(&pageForm)
	list, total, err := models.Picture{}.GetByPage(
		&pageForm.Pagination,
		pageForm.Key,
		pageForm.StorageName,
	)
	if err != nil {
		log.Logger.Sugar().Error("err: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "查询成功",
		Data: utils.GetPageData(list, total, pageForm.Pagination),
	})
}

// UploadImgToAttachment
// @Summary 上传图片到附件
// @Tags 图床
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/images/attachment/upload [post]
func (p *PictureHandler) UploadImgToAttachment(ctx *gin.Context) {
	multiForm, _ := ctx.MultipartForm()
	files := multiForm.File["file[]"]

	if len(files) > 5 || len(files) == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "上传文件个数不能少于 1 个，也不能多于 5 个",
			Data: nil,
		})
		return
	}

	for _, file := range files {
		if !utils.IsImageFile(file.Filename) {
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.RequestError,
				Msg:  "只支持上传 jpeg, jpg, png, gif, bmp 格式的图片",
				Data: nil,
			})
			return
		}
		if !utils.CheckFileSize(file.Size, 5*1024*1024) {
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.RequestError,
				Msg:  "文件大小不能超过 5 MB",
				Data: nil,
			})
			return
		}
	}

	picBedSetting, err := models.SysSettingItem{}.GetBySysSettingName("图床设置")
	if (err != nil && err == gorm.ErrRecordNotFound) || len(picBedSetting) == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "请先配置图床",
			Data: nil,
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	imgSetting, err := models.SysSettingItem{}.GetBySysSettingName(picBedSetting["storage_type"])
	if (err != nil && err == gorm.ErrRecordNotFound) || len(imgSetting) == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "请先配置图床",
			Data: nil,
		})
		return
	}
	if err != nil {
		log.Logger.Sugar().Error("err: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}

	imgUrl := ""
	for _, file := range files {
		switch picBedSetting["storage_type"] {
		case "7bu":
			imgUrl, err = uploadToQubu(file, imgSetting["token"])
		case "sm.ms":
			imgUrl, err = uploadToSmms(file, imgSetting["token"])
		case "imgbb":
			imgUrl, err = uploadToImgbb(file, imgSetting["token"])
		case "cos":
			filePath, fileName, err := saveFile(file, ctx)
			if err != nil {
				log.Logger.Sugar().Error("err: ", err.Error())
				ctx.JSON(http.StatusOK, utils.Result{
					Code: utils.ServerError,
					Msg:  "服务器端错误",
					Data: nil,
				})
				return
			}
			imgUrl, err = uploadToTencentCOS(filePath, fileName, file.Size, imgSetting)
		}
		if err != nil {
			log.Logger.Sugar().Error("图片上传失败：", err.Error())
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.ServerError,
				Msg:  "服务器端错误",
				Data: nil,
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "上传成功",
		Data: gin.H{"imgUrl": imgUrl},
	})
}

// MultiDelPictures
// @Summary 批量删除图片
// @Tags 图床
// @version 1.0
// @Accept application/json
// @Param ids query string true "ids"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/images [delete]
func (p *PictureHandler) MultiDelPictures(ctx *gin.Context) {
	ids := ctx.Query("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请选择要删除的图片",
			Data: nil,
		})
		return
	}
	err := models.Picture{}.MultiDelByIds(ids)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器内部错误",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}

// 上传图片到去不图床
func uploadToQubu(file *multipart.FileHeader, token string) (string, error) {
	// 读取文件
	src, err := file.Open()
	if err != nil {
		log.Logger.Sugar().Error("打开文件失败: ", err.Error())
		return "", err
	}
	defer src.Close()

	// 封装请求体，发送 post 请求并解析响应
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	formFile, _ := bodyWriter.CreateFormFile("file", file.Filename)
	_, _ = io.Copy(formFile, src)
	_ = bodyWriter.Close() // 发送之前必须调用 Close() 以写入结尾行

	req, err := http.NewRequest("POST", qubuUploadURL, bodyBuf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resultMap := map[string]interface{}{}
	err = json.Unmarshal(result, &resultMap)
	if err != nil {
		return "", err
	}

	if !resultMap["status"].(bool) {
		return "", errFileUpload
	}

	data := resultMap["data"].(map[string]interface{})
	links := data["links"].(map[string]interface{})
	sizeStr := data["size"].(string)
	size, err := strconv.ParseFloat(sizeStr, 64)
	if err != nil {
		return "", err
	}
	picture := models.Picture{
		StorageType: "7bu",
		Hash:        data["key"].(string),
		FileName:    data["name"].(string),
		URL:         links["url"].(string),
		Size:        uint(size),
	}
	if err = picture.Create(); err != nil {
		return "", err
	}

	return picture.URL, nil
}

// 上传图片到 sm.ms
func uploadToSmms(file *multipart.FileHeader, token string) (string, error) {
	// 读取文件
	src, err := file.Open()
	if err != nil {
		log.Logger.Sugar().Error("打开文件失败: ", err.Error())
		return "", err
	}
	defer src.Close()

	// 封装请求体，发送 post 请求并解析响应
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	formFile, _ := bodyWriter.CreateFormFile("smfile", file.Filename)
	_, _ = io.Copy(formFile, src)
	_ = bodyWriter.Close() // 发送之前必须调用 Close() 以写入结尾行

	req, err := http.NewRequest("POST", smmsUploadURL, bodyBuf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resultMap := map[string]interface{}{}
	err = json.Unmarshal(result, &resultMap)
	if err != nil {
		return "", err
	}

	if !resultMap["success"].(bool) {
		return "", errFileUpload
	}
	data := resultMap["data"].(map[string]interface{})
	picture := models.Picture{
		StorageType: "sm.ms",
		Hash:        data["hash"].(string),
		FileName:    data["storename"].(string),
		URL:         data["url"].(string),
		Size:        uint(data["size"].(float64)) / 1024,
	}
	if err = picture.Create(); err != nil {
		return "", err
	}

	return picture.URL, nil
}

// 上传图片到 Imgbb
func uploadToImgbb(file *multipart.FileHeader, token string) (string, error) {
	// 读取文件，将其转换成 base64 格式
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return "", err
	}
	base64data := base64.StdEncoding.EncodeToString(fileBytes)

	// 封装请求体，发送 post 请求并解析响应
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	err = bodyWriter.WriteField("image", base64data)
	if err != nil {
		return "", err
	}
	bodyWriter.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", imgbbUploadURL, token), bodyBuf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resultMap := map[string]interface{}{}
	err = json.Unmarshal(result, &resultMap)
	if err != nil {
		return "", err
	}

	if !resultMap["success"].(bool) {
		return "", errFileUpload
	}

	//log.Logger.Sugar().Info("resultMap: ", resultMap)
	data := resultMap["data"].(map[string]interface{})
	picture := models.Picture{
		StorageType: "imgbb",
		Hash:        data["title"].(string),
		FileName:    data["image"].(map[string]interface{})["filename"].(string),
		URL:         data["url"].(string),
		Size:        uint(data["size"].(float64)) / 1024,
	}

	if err = picture.Create(); err != nil {
		return "", err
	}

	return picture.URL, nil
}

// 上传图片到腾讯云 COS
func uploadToTencentCOS(filePath, fileName string, size int64, cosSetting map[string]string) (string, error) {
	u, _ := url.Parse(cosSetting["scheme"] + "://" + cosSetting["host"])
	b := &cos.BaseURL{BucketURL: u}

	// 创建 client
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosSetting["secret_id"],
			SecretKey: cosSetting["secret_key"],
		},
	})
	if client != nil {
		// 上传图片
		result, _, err := client.Object.Upload(
			context.Background(), cosSetting["folder_path"]+"/"+fileName, filePath, nil,
		)
		if err != nil {
			return "", err
		}
		// 删除本地图片
		_ = os.Remove(filePath)

		// 保存图片
		picture := models.Picture{
			StorageType: "cos",
			Hash:        "",
			FileName:    fileName,
			URL:         cosSetting["scheme"] + "://" + result.Location + cosSetting["img_process"],
			Size:        uint(size / 1024),
		}
		if err = picture.Create(); err != nil {
			return "", err
		}

		return picture.URL, nil
	}

	return "", errCosClientCreate
}

// 保存文件到本地
func saveFile(file *multipart.FileHeader, ctx *gin.Context) (filePath string, fileName string, err error) {
	home, _ := utils.Home()
	dirPath := filepath.Join(home, "aries", "images")

	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return
	}

	fileType := utils.GetFileSuffix(file.Filename)
	fileName = strconv.FormatInt(time.Now().Unix(), 10) + fileType // 根据时间戳生成文件名，防止重名
	filePath = filepath.Join(dirPath, fileName)
	err = ctx.SaveUploadedFile(file, filePath)

	return
}
