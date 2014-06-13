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

// 是不是空字符
func IsSpace(c byte) bool {
	if c >= 0x00 && c <= 0x20 {
		return true
	}
	return false
}

// 去掉一个字符串左右的空白串，即（0x00 - 0x20 之内的字符均为空白字符）
// 与strings.TrimSpace功能一致
func Trim(s string) string {
	size := len(s)
	if size <= 0 {
		return s
	}
	l := 0
	for ; l < size; l++ {
		b := s[l]
		if !IsSpace(b) {
			break
		}
	}
	r := size - 1
	for ; r >= l; r-- {
		b := s[r]
		if !IsSpace(b) {
			break
		}
	}
	return string(s[l : r+1])
}

// 去掉一个字符串左右的空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func TrimBytes(bs []byte) string {
	r := len(bs) - 1
	if r <= 0 {
		return string(bs)
	}
	l := 0
	for ; l <= r; l++ {
		b := bs[l]
		if !IsSpace(b) {
			break
		}
	}
	for ; r >= l; r-- {
		b := bs[r]
		if !IsSpace(b) {
			break
		}
	}
	return string(bs[l : r+1])
}

// Trim并且去掉中间多余的空白(多个空白变一个空白)
// 比如 " a b  c    d e" -> "a b c d e"
func TrimExtraSpace(s string) string {
	s = Trim(s)
	size := len(s)
	switch size {
	case 0, 1, 2, 3:
		return s
	default:
		bs := make([]byte, 0, size)
		isSpace := false
		for i := 0; i < size; i++ {
			c := s[i]
			if !IsSpace(c) {
				if isSpace {
					bs = append(bs, ' ')
					isSpace = false
				}
				bs = append(bs, c)
			} else {
				if !isSpace {
					isSpace = true
				}
			}
		}
		return string(bs)
	}
	// 兼容低版本GO
	return ""
}
