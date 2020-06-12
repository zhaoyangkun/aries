package util

import (
	"aries/config/setting"
	"github.com/go-playground/validator/v10"
	"log"
)

// 获取表单错误
func GetFormError(err error) string {
	errMap := err.(validator.ValidationErrors).Translate(setting.Trans) // 将英文表单错误信息翻译，并获取错误信息字典
	var keyList []string                                                // 保存键值
	for key := range errMap {                                           // 遍历 errMap
		keyList = append(keyList, key) // 将 errMap 中的键值保存到 keyList 中
	}
	log.Println(errMap)
	if len(keyList) > 0 {
		return errMap[keyList[0]] // 返回字典中第一个错误信息
	}
	return ""
}

// 判断字符串数组是否包含某个元素
func IsContain(itemList []string, item string) bool {
	for _, val := range itemList {
		if item == val {
			return true
		}
	}
	return false
}
