package localpackage

import "time"

// 获取当前时间
func GetCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format(time.DateTime)
}
