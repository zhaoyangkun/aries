package utils

import "html/template"

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
