package test

import (
	"aries/models"
	"aries/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestReadThemeYaml(t *testing.T) {
	rootPath, _ := os.Getwd()
	dirName := filepath.Join(rootPath, "resources", "themes")
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println("err: ", err.Error())
	}

	yamlPath := ""
	themeInfo := &models.ThemeInfo{}
	for _, file := range fileInfos {
		yamlPath = filepath.Join(dirName, file.Name(), "theme.yaml")
		if utils.FileIsExists(yamlPath) {
			yamlFile, err := ioutil.ReadFile(yamlPath)
			if err != nil {
				log.Panicln("读取配置文件失败：", err.Error())
			}

			err = yaml.Unmarshal(yamlFile, themeInfo)
			if err != nil {
				log.Panicln("配置参数转换失败：", err.Error())
			}
			log.Println(themeInfo)
		}
	}
}
