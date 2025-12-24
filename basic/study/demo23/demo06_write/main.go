package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写文件

	writer := bufio.NewWriter(file)
	//writer.WriteString("你好golang")

	for i := 0; i < 10; i++ {
		writer.WriteString("直接写入的字符串数据" + strconv.Itoa(i) + "\r\n")
	}

	writer.Flush()
}
