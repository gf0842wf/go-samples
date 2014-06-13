package ztime

import (
	"strings"
	"time"
)

func NowUnix() int64 {
	return time.Now().Unix()
}

func fmt2fmt(fmt string) string {
	// "%Y-%m-%d %H:%M:%S" = "%Y-%m-%d %X"
	if fmt == "" { // 默认情况
		fmt = "2006-01-02 15:04:05"
	} else if strings.Count(fmt, "%") != 0 {
		fmt = strings.Replace(fmt, "%Y", "2006", 1)
		fmt = strings.Replace(fmt, "%m", "01", 1)
		fmt = strings.Replace(fmt, "%d", "02", 1)
		fmt = strings.Replace(fmt, "%H", "15", 1)
		fmt = strings.Replace(fmt, "%M", "04", 1)
		fmt = strings.Replace(fmt, "%S", "05", 1)
		fmt = strings.Replace(fmt, "%X", "15:04:05", 1)
	}

	return fmt
}

func NowFmt(fmt string) string {
	fmt = fmt2fmt(fmt)

	return time.Now().Format(fmt)
}

// 下面两个是转化为time.Time, 然后就可以加减,取具体年月日星期等
func Unix2Struct(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}

func Fmt2Struct(fmt string, fmtTime string) time.Time {
	fmt = fmt2fmt(fmt)
	theTime, _ := time.Parse(fmt, fmtTime)
	return theTime
}

func Unix2Fmt(fmt string, unixTime int64) string {
	fmt = fmt2fmt(fmt)
	return time.Unix(unixTime, 0).Format(fmt)
}

func Fmt2Unix(fmt string, fmtTime string) int64 {
	return Fmt2Struct(fmt, fmtTime).Unix()
}
