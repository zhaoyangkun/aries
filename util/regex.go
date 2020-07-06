package util

import "regexp"

// 获取 HTML 标签中的纯文本
func GetHtmlContent(html string) (content string) {
	reg := regexp.MustCompile(`<[\S\s]+?>`)
	content = reg.ReplaceAllString(html, "")
	return
}
