package main

import (
	"fmt"
	"time"
)

// 打印当前时间
func main() {
	timeObj := time.Now()
	fmt.Println(timeObj) //2025-02-20 17:30:45.9519865 +0800 CST m=+0.003181601

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) //2025-02-20 17:30:45
	fmt.Println("\n----------------\n")
	fmt.Printf(timeObj.Format("2006-01-02 15:04:05")) //15 代表小时数，24小时制
	fmt.Println("\n----------------\n")
	fmt.Printf(timeObj.Format("2006-01-02 03:04:05")) // 03 代表小时数，12小时制
}
