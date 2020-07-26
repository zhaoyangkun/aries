package setting

import (
	"github.com/88250/lute"
	ut "github.com/go-playground/universal-translator"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 总配置
type Setting struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"db"`
}

// 项目配置结构
type server struct {
	Mode            string   `yaml:"mode"`              // 运行模式
	Port            string   `yaml:"port"`              // 运行端口
	TokenExpireTime int      `yaml:"token_expire_time"` // JWT token 过期时间
	AllowedRefers   []string `yaml:"allowed_refers"`    // 允许的 referer
}

// 数据库配置结构
type database struct {
	Host        string `yaml:"host"`          // 主机地址
	UserName    string `yaml:"user_name"`     // 用户名
	Password    string `yaml:"password"`      // 密码
	Database    string `yaml:"database"`      // 数据库名
	Port        string `yaml:"port"`          // 端口
	TimeZone    string `yaml:"time_zone"`     // 时区
	MaxIdleConn int    `yaml:"max_idle_conn"` // 最大空闲连接数
	MaxOpenConn int    `yaml:"max_open_conn"` // 最大打开连接数
}

// 翻译器
var Trans ut.Translator

// 全局配置
var Config = &Setting{}

// lute
var LuteEngine = &lute.Lute{}

// 读取 yaml 配置文件
func InitSetting() {
	// 获取当前项目根目录
	rootPath, _ := os.Getwd()
	// 解决 GoLand 默认单元测试环境下，读取配置文件失败的问题
	rootPath = strings.Replace(rootPath, "test", "", -1)
	// 拼接配置文件访问路径
	yamlPath := filepath.Join(rootPath, "config", "develop.yaml")
	log.Println("yamlPath: ", yamlPath)
	// 读取配置文件
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Panicln("读取配置文件失败：", err.Error())
	}
	// 转换配置文件参数
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Panicln("配置参数转换失败：", err.Error())
	}
}
