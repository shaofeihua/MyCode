package utils

// 检查字符串是否只包含英文字母和数字（也可以用正则表达式来检查，但是用字符串比较的方式比正则的运行效率更高）
func IsLetterOrNum(txt string) bool {
	for _, ch := range txt {
		if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
			return false
		}
	}
	return true
}
