package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	smmsUploadURL  = "https://sm.ms/api/v2/upload"
	imgbbUploadURL = "https://api.imgbb.com/1/upload"
)

var imgTypes = []string{".jpeg", ".jpg", ".png", ".gif", ".bmp"}

// 获取文件名
func GetFileNameOnly(filePath string) (fileName string) {
	fileNameWithSuffix := path.Base(filePath)                     // 获取文件名带后缀
	fileSuffix := path.Ext(fileNameWithSuffix)                    // 获取文件后缀
	fileName = strings.TrimSuffix(fileNameWithSuffix, fileSuffix) // 获取文件名
	return
}

// 判断文件是否存在
func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 判断是否为图片格式
func IsImageFile(fileName string) bool {
	fileSuffix := path.Ext(fileName)
	for _, val := range imgTypes {
		if val == fileSuffix {
			return true
		}
	}
	return false
}

// 上传图片到 sm.ms
func UploadToSmms(file *multipart.FileHeader, token string) error {
	// 读取文件
	src, err := file.Open()
	if err != nil {
		log.Errorln("打开文件失败: ", err.Error())
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
		log.Errorln("创建请求失败: ", err.Error())
		return err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorln("发送请求失败: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln("读取响应内容失败: ", err.Error())
		return err
	}
	log.Infoln("resp: ", string(result))
	return nil
}

// 上传图片到 Imgbb
func UploadToImgbb(file *multipart.FileHeader, token string) error {
	// 读取文件，将其转换成 base64 格式
	src, err := file.Open()
	if err != nil {
		log.Errorln("打开文件失败: ", err.Error())
		return err
	}
	defer src.Close()
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		log.Errorln("读取文件内容失败: ", err.Error())
		return err
	}
	base64data := base64.StdEncoding.EncodeToString(fileBytes)

	// 封装请求体，发送 post 请求并解析响应
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	err = bodyWriter.WriteField("image", base64data)
	if err != nil {
		log.Errorln("封装请求失败: ", err.Error())
		return err
	}
	bodyWriter.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", imgbbUploadURL, token), bodyBuf)
	if err != nil {
		log.Errorln("创建请求失败: ", err.Error())
		return err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorln("发送请求失败: ", err.Error())
		return err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln("读取响应内容失败: ", err.Error())
		return err
	}
	log.Infoln("resp: ", string(result))
	return nil
}
