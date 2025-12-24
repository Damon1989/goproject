package goselect

import (
	"fmt"
	"time"
)

func SelectSmt() {
	var a [4]int
	var c1, c2, c3, c4 = make(chan int), make(chan int), make(chan int), make(chan int)
	var i1, i2 = 0, 42

	go func() {
		c1 <- 10
	}()

	go func() {
		<-c2
	}()

	go func() {
		close(c3)
	}()

	go func() {
		c4 <- 40
	}()

	go func() {
		select {
		case i1 = <-c1:
			fmt.Println("received", i1, " from c1")
		case c2 <- i2:
			fmt.Println("sent", i2, " to c2")
		case i3, ok := <-c3:
			if ok {
				fmt.Println("received", i3, " from c3")
			} else {
				fmt.Println("c3 is closed")
			}
		case a[f()] = <-c4:
			fmt.Println("received", a[f()], " from c4")
		default:
			fmt.Println("no communication")
		}
	}()

	fmt.Println("select")
	time.Sleep(10 * time.Second)
}

func f() int {
	print("f()被调用\n")
	return 2
}

func SelectMain() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		stringChan <- fmt.Sprintf("str%d", i)
	}

	// 使用select来获取channel里面的数据的时候不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Println("Received from intChan:", v)
		case v := <-stringChan:
			fmt.Println("Received from stringChan:", v)
		default:
			fmt.Println("No more data to receive. Exiting.")
			return
		}
	}
}
