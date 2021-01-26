package utils

import (
	"html/template"
	"time"
)

// 取消 html 转义
func SafeHtml(x string) interface{} {
	return template.HTML(x)
}

// 两数相加
func AddUpTwoNum(a, b int) int {
	return a + b
}

// 两数相减
func SubtractTwoNum(a, b int) int {
	return a - b
}

// 求余
func Mod(a, b int) int {
	return a % b
}

// 获取年份
func Year(datetime time.Time) int {
	return datetime.Year()
}

// 获取月份
func Month(datetime time.Time) int {
	return int(datetime.Month())
}

// 获取日
func Day(datetime time.Time) int {
	return int(datetime.Day())
}
