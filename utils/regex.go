package utils

import "regexp"

// GetHtmlContent 获取 HTML 标签中的纯文本
func GetHtmlContent(html string) (content string) {
	reg := regexp.MustCompile(`(<[\S\s]+?>)|([\s]+?)`)
	content = reg.ReplaceAllString(html, "")

	return
}
