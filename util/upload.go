package util

import (
	"path"
	"strings"
)

func GetFileNameOnly(filePath string) (fileName string) {
	fileNameWithSuffix := path.Base(filePath)                     // 获取文件名带后缀
	fileSuffix := path.Ext(fileNameWithSuffix)                    // 获取文件后缀
	fileName = strings.TrimSuffix(fileNameWithSuffix, fileSuffix) // 获取文件名
	return
}
