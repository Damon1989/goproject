package main

import "fmt"

func main() {
	//var ch1 = make(chan int, 10)
	//for i := 0; i < 10; i++ {
	//	ch1 <- i
	//}
	//
	//close(ch1) // 关闭管道
	//// for range 循环遍历管道的值，注意：管道没有key
	//for v := range ch1 {
	//	fmt.Println(v)
	//}

	var ch2 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch2 <- i
	}

	for i := 0; i < 10; i++ {
		v := <-ch2
		fmt.Println(v)
	}
}
