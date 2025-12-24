package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./mainold.go")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	// bufio读取文件
	reader := bufio.NewReader(file)
	var fileStr string
	for {
		str, err := reader.ReadString('\n') // 表示一次读取一行
		if err == io.EOF {
			fileStr += str
			break
		}
		for err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}

	fmt.Println(fileStr)
}
