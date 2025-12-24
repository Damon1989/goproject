package main

import (
	"fmt"
)

func main() {

	//1、创建channel
	ch := make(chan int, 3)
	//2、给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	//3、获取管道里面的内容
	a := <-ch
	fmt.Println(a)
	<-ch // 从管道里面取值，但是不赋值给任何变量

	c := <-ch
	fmt.Println(c)

	ch <- 56
	ch <- 66

	//4、打印管道的长度和容量
	fmt.Printf("值：%v 容量：%v 长度：%v\n", ch, cap(ch), len(ch))

	//5、管道的类型（引用数据类型）
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 69

	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d) //69

	//6、管道阻塞
	ch6 := make(chan int, 1)

	ch6 <- 34
	//ch6 <- 34 // all goroutines are asleep - deadlock!

	ch7 := make(chan string, 2)

	ch7 <- "数据1"
	ch7 <- "数据2"

	m1 := <-ch7
	m2 := <-ch7
	//m3 := <-ch7 // all goroutines are asleep - deadlock!
	fmt.Println(m1, m2)

	fmt.Println("----------------")
	ch8 := make(chan int, 1)
	ch8 <- 34
	<-ch8
	ch8 <- 341
	<-ch8
	ch8 <- 342
	m4 := <-ch8
	fmt.Println(m4)

}
