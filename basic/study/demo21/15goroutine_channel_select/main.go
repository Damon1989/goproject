package main

import (
	"fmt"
	"time"
)

// 在某些场景下我们需要同事从多个通道接受数据，这个时候我们就可以用到golang中给我们提供的select多路复用功能。
func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println("intChan:", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Println("stringChan:", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("数据获取完毕")
			return
		}
	}
}
