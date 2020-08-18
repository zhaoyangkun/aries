package api

import (
	"aries/log"
	"aries/models"
	"aries/utils"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	smmsUploadURL  = "https://sm.ms/api/v2/upload"
	imgbbUploadURL = "https://api.imgbb.com/1/upload"
)

type PictureHandler struct {
}

// @Summary 上传图片到附件
// @Tags图床
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/img/upload [post]
func (p *PictureHandler) UploadImgToAttachment(ctx *gin.Context) {
	multiForm, _ := ctx.MultipartForm()
	files := multiForm.File["file[]"]
	for _, file := range files {
		if !utils.IsImageFile(file.Filename) {
			ctx.JSON(http.StatusOK, utils.Result{
				Code: utils.RequestError,
				Msg:  "只支持上传jpeg, jpg, png, gif, bmp 格式的图片",
				Data: nil,
			})
			return
		}
	}
	picBedSetting, err := models.SysSettingItem{}.GetBySysSettingName("图床设置")
	if err != nil {
		log.Logger.Sugar().Error("err: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	if len(picBedSetting) == 0 {
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.RequestError,
			Msg:  "请先配置图床",
			Data: nil,
		})
		return
	}
	imgSetting, err := models.SysSettingItem{}.GetBySysSettingName(picBedSetting["storage_type"])
	if err != nil {
		log.Logger.Sugar().Error("err: ", err.Error())
		ctx.JSON(http.StatusOK, utils.Result{
			Code: utils.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	for _, file := range files {
		switch picBedSetting["storage_type"] {
		case "sm.ms":
			err = UploadToSmms(file, imgSetting["token"])
		case "imgbb":
			err = UploadToImgbb(file, imgSetting["token"])
		case "cos":
			break
		}
		if err != nil {
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
		Data: nil,
	})
}

// 上传图片到 sm.ms
func UploadToSmms(file *multipart.FileHeader, token string) error {
	// 读取文件
	src, err := file.Open()
	if err != nil {
		log.Logger.Sugar().Error("打开文件失败: ", err.Error())
		return err
	}
	defer src.Close()

	// 封装请求体，发送 post 请求并解析响应
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	formFile, _ := bodyWriter.CreateFormFile("smfile", file.Filename)
	_, _ = io.Copy(formFile, src)
	bodyWriter.Close() // 发送之前必须调用Close()以写入结尾行

	req, err := http.NewRequest("POST", smmsUploadURL, bodyBuf)
	if err != nil {
		log.Logger.Sugar().Error("创建请求失败: ", err.Error())
		return err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Logger.Sugar().Error("发送请求失败: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Sugar().Error("读取响应内容失败: ", err.Error())
		return err
	}
	log.Logger.Sugar().Error("resp: ", string(result))
	return nil
}

// 上传图片到 Imgbb
func UploadToImgbb(file *multipart.FileHeader, token string) error {
	// 读取文件，将其转换成 base64 格式
	src, err := file.Open()
	if err != nil {
		log.Logger.Sugar().Error("打开文件失败: ", err.Error())
		return err
	}
	defer src.Close()
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		log.Logger.Sugar().Error("读取文件内容失败: ", err.Error())
		return err
	}
	base64data := base64.StdEncoding.EncodeToString(fileBytes)

	// 封装请求体，发送 post 请求并解析响应
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	err = bodyWriter.WriteField("image", base64data)
	if err != nil {
		log.Logger.Sugar().Error("封装请求失败: ", err.Error())
		return err
	}
	bodyWriter.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", imgbbUploadURL, token), bodyBuf)
	if err != nil {
		log.Logger.Sugar().Error("创建请求失败: ", err.Error())
		return err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Logger.Sugar().Error("发送请求失败: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Sugar().Error("读取响应内容失败: ", err.Error())
		return err
	}
	log.Logger.Sugar().Error("resp: ", string(result))
	return nil
}
