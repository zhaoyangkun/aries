package utils

import (
	"aries/config/setting"
	"html/template"
	"time"
)

// SafeHtml 取消 html 转义
func SafeHtml(x string) interface{} {
	return template.HTML(x)
}

// AddUpTwoNum 两数相加
func AddUpTwoNum(a, b int) int {
	return a + b
}

// SubtractTwoNum 两数相减
func SubtractTwoNum(a, b int) int {
	return a - b
}

// Mod 求余
func Mod(a, b int) int {
	return a % b
}

// Year 获取年份
func Year(datetime time.Time) int {
	return datetime.Year()
}

// Month 获取月份
func Month(datetime time.Time) int {
	return int(datetime.Month())
}

// Day 获取日
func Day(datetime time.Time) int {
	return int(datetime.Day())
}

// GetTheme 获取主题
func GetTheme() string {
	return setting.BlogVars.Theme + "/"
}
