package utils

import (
	"os"
	"path"
	"strings"
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
