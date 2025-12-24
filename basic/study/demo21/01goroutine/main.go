package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*

	在主线程（可以理解成进程）中,开启一个goroutine,该协程每隔50毫秒输出“你好golang”
	在主线程中也每隔50毫秒输出“你好world” ，然后等待两个协程执行完毕后，退出主线程。
	要求主线程和协程同时执行
*/
var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好golang", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程执行完毕，协程数量-1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2() 你好golang", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程执行完毕，协程数量-1
}

func main() {
	wg.Add(1)  // 协程数量+1
	go test1() //开启一个协程
	wg.Add(1)  // 协程数量+1
	go test2()
	for i := 0; i < 10; i++ {
		fmt.Println("main() 你好golang", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
	fmt.Println("主线程退出...")
}
