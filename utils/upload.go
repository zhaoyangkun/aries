package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
)

var imgTypes = []string{".jpeg", ".jpg", ".png", ".gif", ".bmp"}

// GetFileNameOnly 获取文件名
func GetFileNameOnly(filePath string) (fileName string) {
	fileNameWithSuffix := path.Base(filePath)                     // 获取文件名带后缀
	fileSuffix := path.Ext(fileNameWithSuffix)                    // 获取文件后缀
	fileName = strings.TrimSuffix(fileNameWithSuffix, fileSuffix) // 获取文件名

	return
}

// FileIsExists 判断文件是否存在
func FileIsExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// IsImageFile 判断是否为图片格式
func IsImageFile(fileName string) bool {
	fileType := GetFileSuffix(fileName)

	for _, val := range imgTypes {
		if val == fileType {
			return true
		}
	}

	return false
}

// CheckFileSize 校验文件大小
func CheckFileSize(size int64, maxSize int64) bool {
	return size <= maxSize
}

// GetFileSuffix 获取文件后缀
func GetFileSuffix(filePath string) string {
	fileSuffix := path.Ext(filePath)
	return fileSuffix
}

// 返回主目录
func Home() (string, error) {
	u, err := user.Current()
	if err == nil {
		return u.HomeDir, nil
	}

	// windows
	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path

	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}

	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
