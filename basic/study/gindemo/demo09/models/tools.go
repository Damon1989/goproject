package models

import (
	"fmt"
	"time"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println("时间戳转日期：", timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳
func TimeToUnix(date string) int64 {
	fmt.Println("日期转时间戳：", date)
	t, _ := time.Parse("2006-01-02 15:04:05", date)
	return t.Unix()
}

func getUnix() int64 {
	return time.Now().Unix()
}

func getDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
