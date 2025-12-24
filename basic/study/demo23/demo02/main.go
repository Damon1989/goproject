package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1.打开文件
	file, err := os.Open("./mainold.go")
	defer file.Close() //必须得关闭文件流
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	//	2.读取文件
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败")
			return
		}
		strSlice = append(strSlice, tempSlice[:n]...) // 注意写法
	}

	fmt.Println(string(strSlice))
}
