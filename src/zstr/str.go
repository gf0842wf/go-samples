package zstr

import (
	"strconv"
	"strings"
)

// Partition(s, sep) -> (head, sep, tail)
func Partition(s string, sep string) (head string, retSep string, tail string) {
	index := strings.Index(s, sep)
	if index == -1 {
		head = s
		retSep = ""
		tail = ""
	} else {
		head = s[:index]
		retSep = sep
		tail = s[len(head)+len(sep):]
	}
	return
}

/***********************

// 包格式 头(xy) + 数据体 + 尾 (..xy...)
// ...
_, header, msg := Partition(data, "xy")
if header == "" {
    // 没有头(xy)丢包. (也有可能粘包分包导致 "...x", 最后一个(注意是一个)字符变成了x, 这时要把前面的包丢弃,只保留一个x)
} else {
    // do
}

***********************/

// 转化并拼接
// ToString("abcd", 12, 45)
// => abcd1245
func ToString(args ...interface{}) string {
	result := ""
	for _, arg := range args {
		switch val := arg.(type) {
		case int:
			result += strconv.Itoa(val)
		case string:
			result += val
		}
	}
	return result
}
