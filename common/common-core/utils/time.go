package utils

import "time"

const (
	// TimeFormat 标准时间格式
	TimeFormat = "2006-01-02 15:04:05"
	// DateFormat 标准日期格式
	DateFormat = "2006-01-02"
)

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format(TimeFormat)
}

// FormatDate 格式化日期
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// ParseTime 解析时间字符串
func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation(TimeFormat, str, time.Local)
}

// ParseDate 解析日期字符串
func ParseDate(str string) (time.Time, error) {
	return time.ParseInLocation(DateFormat, str, time.Local)
}

// Now 获取当前时间
func Now() time.Time {
	return time.Now()
}

// NowUnix 获取当前时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}


