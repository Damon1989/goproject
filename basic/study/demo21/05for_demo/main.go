package main

import (
	"fmt"
	"time"
)

// 需求：要统计1-100的数字中哪些是素数？for循环
// 1. 遍
func main() {
	start := time.Now().UnixMilli()
	for num := 2; num < 120000; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%d是素数\n", num)
		}
	}
	end := time.Now().UnixMilli()
	fmt.Println("耗时：", end-start)

}
