package valiUtil

//验证规则相关

import (
	"errors"
	"unicode"

	maputil "github.com/lily0749labs/goutils/maps"
	validutil "github.com/lily0749labs/goutils/valid"
)

type Rules struct {
	Mkey  string
	Value string
}

// 验证参数是否存在
func ValidateData(v_data map[string]any, rules []Rules) (bool, string) {
	for _, value := range rules {
		param := value.Mkey
		rule := value.Value
		if !maputil.ContainsKeys(v_data, param) {
			return false, rule
		}
		if v_data[param] == nil {
			return false, rule
		}
	}
	return true, ""
}

/*
*验证电话号码是否正确
 */
func ValidatePhone(phone string) bool {
	return validutil.Valid.Mobile(phone) || validutil.Valid.Telephone(phone)
}

/*
*验证身份证是否正确
 */
func ValidateIdCard(idCard string) (bool, error) {
	if len(idCard) != 18 {
		return false, errors.New("必须要输入18位的身份证号码")
	}
	if !validutil.Valid.IDCard18(idCard) {
		return false, errors.New("身份证号码不正确")
	}
	return true, nil
}

/*
判断是否为符号
*/
func isSymbol(str string) bool {
	res := false
	symbol_array := [...]string{"~", "`", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "\"", ":", ";", "*", "/", ".", "=", "<", ",", ">", "?", "\\", "/", "|", "{", "}", "[", "]"}
I:
	for i := 0; i < len(symbol_array); i++ {
		if str == symbol_array[i] {
			res = true
			break I
		}
	}
	return res
}

/*
验证用户名规则
验证通过为true，否则为false
*/
func VerifyUserName(str string) bool {
	res := true
I:
	for _, r := range str {
		if unicode.IsLetter(r) { // 判断是否为字母
		} else if unicode.IsDigit(r) { // 判断是否为数字
		} else if isSymbol(string(r)) { // 判断是否为符号

		} else {
			res = false
			break I
		}
	}
	return res
}
