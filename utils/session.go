package utils

import "github.com/gin-contrib/sessions"

// GetSessionStr 根据 key 获取 string 类型的 session 值
func GetSessionStr(session sessions.Session, key string) string {
	value := session.Get(key)
	if value == nil {
		value = ""
	}

	valueStr := value.(string) //将interface{}转为string

	return valueStr
}
