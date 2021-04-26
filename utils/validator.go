package utils

import (
	"aries/config/setting"
	"strings"

	"github.com/go-playground/validator/v10"
)

var transMap = map[string]string{
	"Pwd": "密码",
}

// 自定义翻译函数
func transTagName(libTans, err interface{}) string {
	errs := make(map[string]string)
	for k, v := range err.(validator.ValidationErrorsTranslations) {
		for key, value := range libTans.(map[string]string) {
			v = strings.Replace(v, key, value, -1)
		}
		errs[k] = v
	}

	var keyList []string    // 保存键值
	for key := range errs { // 遍历 errMap
		keyList = append(keyList, key) // 将 errMap 中的键值保存到 keyList 中
	}

	if len(keyList) > 0 {
		return errs[keyList[0]] // 返回字典中第一个错误信息
	}

	return ""
}

// GetFormError 获取表单错误
func GetFormError(err error) string {
	// 翻译表单错误
	if errs, ok := err.(validator.ValidationErrors); ok {
		return transTagName(transMap, errs.Translate(setting.Trans))
	}

	return "数据类型转换错误"
}

// IsContain 判断字符串数组是否包含某个元素
func IsContain(itemList []string, item string) bool {
	for _, val := range itemList {
		if item == val {
			return true
		}
	}

	return false
}
