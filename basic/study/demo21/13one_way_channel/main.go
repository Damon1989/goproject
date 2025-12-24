package main

import "fmt"

// 单向管道
func main() {
	//1、在默认情况下，管道是双向的
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 12
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2)
	//2、管道声明为只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12
	//m3 := <-ch2 //invalid operation: <-ch2 (receive from send-only type chan<- int)
	//3、管道声明为只读
	ch3 := make(<-chan int, 2)
	//ch3 <- 10 //invalid operation: ch3 <- 10 (send to receive-only type <-chan int)
	m4 := <-ch3
	m5 := <-ch3
	fmt.Println(m4, m5)

}
