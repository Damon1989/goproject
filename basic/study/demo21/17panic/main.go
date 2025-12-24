package main

import (
	"fmt"
	"time"
)

// 函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello,world")
	}
}

// 函数

func test() {
	// 这里我们可以使用defer + recover
	defer func() {
		// recover() 内置函数，可以捕获到异常
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[1] = "hello"
}
func main() {
	go sayHello()
	go test()

	// 防止主进程退出这里使用time.Sleep演示，搭建也可以是使用sync.WaitGroup
	time.Sleep(time.Second)
}
