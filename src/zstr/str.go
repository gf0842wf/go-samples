package zstr

import (
	"strings"
)

func Partition(s string, sep string) (head string, retSep string, tail string) {
	// Partition(s, sep) -> (head, sep, tail)
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
