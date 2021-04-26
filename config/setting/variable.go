package setting

import (
	"github.com/88250/lute"
	"github.com/gin-contrib/cache/persistence"
	ut "github.com/go-playground/universal-translator"
)

// BlogVars 博客全局变量
var BlogVars BlogVariable

// Trans 翻译器
var Trans ut.Translator

// Config 全局配置
var Config = &Setting{}

// LuteEngine markdown 引擎
var LuteEngine = &lute.Lute{}

// Cache 缓存
var Cache = &persistence.InMemoryStore{}
