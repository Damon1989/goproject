package main

import "errors"

// 模拟了一个读取文件的方法
func readFile(fileName string) error {
	if fileName == "mainold.go" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			println("给管理员发送邮件", err)
		}
	}()
	err := readFile("***.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	myFn()
	println("结束")
}
