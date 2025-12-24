package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFileName, destFileName string) (err error) {
	sFile, err1 := os.Open(srcFileName)
	dFile, err2 := os.OpenFile(destFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer sFile.Close()
	defer dFile.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	var tempSlice = make([]byte, 128)
	for {
		// 读取数据
		n1, err := sFile.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// 写入数据
		if _, err := dFile.Write(tempSlice[:n1]); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := copyFile("./test.txt", "./test2.txt"); err != nil {
		fmt.Printf("拷贝错误 err=%vn", err)
	} else {
		fmt.Println("拷贝成功")
	}

}
