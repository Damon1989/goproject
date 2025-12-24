package main

import (
	"fmt"
	"io/ioutil"
)

func copy(srcFileName, destFIleName string) (err error) {
	byteStr, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(destFIleName, byteStr, 006)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func main() {
	copy("./test.txt", "./test1.txt")
	fmt.Println("复制文件成功")
}
