package split_string

import "strings"

// 切割字符串
func Split(str string, sep string) []string {
	// Spit 切割字符串
	var ret []string
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
