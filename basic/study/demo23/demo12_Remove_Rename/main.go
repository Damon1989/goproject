package main

import (
	"fmt"
	"os"
)

func main() {
	// 删除一个文件
	//err := os.Remove("./aaa.txt")
	//if err != nil {
	//	fmt.Println("删除文件失败", err)
	//	return
	//}

	// 删除一个目录
	//err := os.Remove("./go")
	//if err != nil {
	//	fmt.Println("删除文件夹失败", err)
	//	return
	//}

	// 一次删除多个目录
	err := os.RemoveAll("./dir1")
	if err != nil {
		fmt.Println("删除文件夹失败", err)
		return
	}

}
