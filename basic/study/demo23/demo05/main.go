package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写文件

	//for i := 0; i < 10; i++ {
	//	file.WriteString("直接写入的字符串数据" + strconv.Itoa(i) + "\r\n")
	//}

	var str = "直接写入的字符串数据"
	file.Write([]byte(str))
}
