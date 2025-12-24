package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	unixtime := timeObj.Unix() // 获取当前时间戳  毫秒
	fmt.Println("当前时间戳：", unixtime)

	unixNano := timeObj.UnixNano()
	fmt.Println("当前时间戳纳秒：", unixNano)

}
