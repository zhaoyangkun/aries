package setting

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/88250/lute"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// BlogVariable 博客全局变量
type BlogVariable struct {
	Theme          string
	ContextPath    string
	StaticRootPath string
	SiteName       string
	SiteDesc       string
	SiteKeywords   string
	SiteLogo       string
	RecordNumber   string
	GlobalHeader   string
	GlobalFooter   string
	QQ             string
	Wechat         string
	Github         string
	Weibo          string
	Zhihu          string
}

// Setting 总配置
type Setting struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"db"`
	Logger   logger   `yaml:"logger"`
	SMTP     smtp     `yaml:"smtp"`
}

// server 服务器配置
type server struct {
	Mode            string   `yaml:"mode"`              // 运行模式
	Port            string   `yaml:"port"`              // 运行端口
	TokenExpireTime int      `yaml:"token_expire_time"` // JWT token 过期时间
	AllowedRefers   []string `yaml:"allowed_refers"`    // 允许的 referer
	LimitTime       int64    `yaml:"limit_time"`        // 限流时间间隔
	LimitCap        int64    `yaml:"limit_cap"`         // 间隔时间内最大访问次数
}

// database 数据库配置
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

// logger 日志
type logger struct {
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
}

// smtp 信息（重置密码 + 评论回复）
type smtp struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}

// InitSetting 读取 yaml 配置文件
func (s *Setting) InitSetting() {
	/* 开发环境 */
	// 获取当前项目根目录
	rootPath, _ := os.Getwd()
	// 解决 GoLand 默认单元测试环境下，读取配置文件失败的问题
	rootPath = strings.Replace(rootPath, "test", "", -1)
	// 拼接配置文件访问路径
	yamlPath := filepath.Join(rootPath, "config", "develop.yaml")

	/* 生产环境 */
	//homeDir, err := os.UserHomeDir()
	//if err != nil {
	//	log.Panicln("获取用户主目录失败：", err.Error())
	//}
	//yamlPath := filepath.Join(homeDir, ".aries", "aries.yaml")

	log.Println("配置文件路径：", yamlPath)
	// 读取配置文件
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Panicln("读取配置文件失败：", err.Error())
	}

	// 转换配置文件参数
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Panicln("配置参数转换失败：", err.Error())
	}
}

// InitLute 初始化 markdown 引擎
func (s *Setting) InitLute() {
	LuteEngine = lute.New()
	LuteEngine.SetCodeSyntaxHighlight(true)
}

// InitCache 初始化缓存
func (s *Setting) InitCache() {
	Cache = persistence.NewInMemoryStore(time.Hour * 1)
}

// InitBlogVars 配置博客全局变量
func (b *BlogVariable) InitBlogVars(siteSetting map[string]string, socialInfo map[string]string) {
	if theme, ok := siteSetting["theme_name"]; ok {
		b.Theme = theme
	} else {
		// 默认主题
		b.Theme = "xue"
		//b.Theme = "boundless-ui"
	}
	if contextPath, ok := siteSetting["site_url"]; ok {
		b.ContextPath = contextPath
	} else {
		b.ContextPath = ""
	}
	// 若用户已经设置静态资源根路径，使用用户设置
	if staticRoot, ok := siteSetting["static_root"]; ok {
		b.StaticRootPath = staticRoot
	} else {
		// 若用户未设置静态资源根路径，在测试环境下使用本地资源路径，生产环境下使用 cdn
		if Config.Server.Mode == gin.DebugMode {
			b.StaticRootPath = b.ContextPath
		} else {
			b.StaticRootPath = "https://gcore.jsdelivr.net/gh/zhaoyangkun/aries@latest"
		}
	}
	if siteName, ok := siteSetting["site_name"]; ok {
		b.SiteName = siteName
	} else {
		b.SiteName = "Aries"
	}
	if siteDesc, ok := siteSetting["site_desc"]; ok {
		b.SiteDesc = siteDesc
	} else {
		b.SiteDesc = "Aries Blog"
	}
	if siteKeywords, ok := siteSetting["seo_key_words"]; ok {
		b.SiteKeywords = siteKeywords
	}
	if siteLogo, ok := siteSetting["site_logo"]; ok {
		b.SiteLogo = siteLogo
	} else {
		b.SiteLogo = "https://s1.ax1x.com/2020/08/07/aWuU41.png"
	}
	if recordNumber, ok := siteSetting["record_number"]; ok {
		b.RecordNumber = recordNumber
	} else {
		b.RecordNumber = ""
	}
	if globalHeader, ok := siteSetting["head_content"]; ok {
		b.GlobalHeader = globalHeader
	}
	if globalFooter, ok := siteSetting["footer_content"]; ok {
		b.GlobalFooter = globalFooter
	}
	if qq, ok := socialInfo["qq"]; ok {
		b.QQ = qq
	}
	if wechat, ok := socialInfo["wechat"]; ok {
		b.Wechat = wechat
	}
	if github, ok := socialInfo["github"]; ok {
		b.Github = github
	}
	if zhihu, ok := socialInfo["zhihu"]; ok {
		b.Zhihu = zhihu
	}
	if weibo, ok := socialInfo["weibo"]; ok {
		b.Weibo = weibo
	}
}
