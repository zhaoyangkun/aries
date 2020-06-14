package setting

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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

// 读取 yaml 配置文件
func init() {
	yamlFile, err := ioutil.ReadFile("config/" + "develop.yaml")
	if err != nil {
		log.Panicln("读取配置文件失败：", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Panicln("配置参数转换失败：", err.Error())
	}
}
