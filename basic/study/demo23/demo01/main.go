package main

import (
	"fmt"
	"os"
)

func main() {
	// 只读方式打开当前目录下的main.go文件 D:\workspace\code\damon\go\goproject\basic\study\demo23\familyaccount\mainold.go
	//file, err := os.Open("D:\\workspace\\code\\damon\\go\\goproject\\basic\\study\\demo23\\familyaccount\\test.txt")
	//file, err := os.Open("./test.txt")
	file, err := os.Open("./mainold.go")
	defer file.Close() //必须得关闭文件流
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	// 操作文件
	fmt.Println(file)
}
