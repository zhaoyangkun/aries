package test

import (
	"aries/models"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

type ThemeYaml struct {
	General General `yaml:"general"`
}

type General struct {
	Label string                      `yaml:"label"`
	Items map[interface{}]interface{} `yaml:"items"`
}

func TestYaml(t *testing.T) {
	theme := ThemeYaml{}
	yamlStr := `
general:
  label: 通用设置
  items:
    home_title:
      name: home_title
      label: 首页展示博客名
      type: text
      placeholder: "用于首页展示名字"
      description: "注意：如填写，会覆盖原有博客名"
    posts_style:
      name: posts_style
      label: 首页文章列表展示风格
      type: radio
      data-type: bool
      default: true
      options:
        - value: true
          label: 卡片
        - value: false
          label: 列表`
	_ = yaml.Unmarshal([]byte(yamlStr), &theme)
	for k, v := range theme.General.Items {
		log.Println("k: ", k)
		log.Println("v: ", v)
	}
}

func TestThemeYaml(t *testing.T) {
	themeInfo := models.ThemeInfo{}
	rootPath, _ := os.Getwd()
	rootPath = strings.Replace(rootPath, "test", "", -1)
	yamlPath := filepath.Join(rootPath, "themes", "xue", "theme.yaml")
	log.Println("主题配置文件路径：", yamlPath)

	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Panicln("读取配置文件失败：", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &themeInfo)
	if err != nil {
		log.Panicln("配置参数转换失败：", err.Error())
	}

	log.Println("themeInfo: ", themeInfo)
}
