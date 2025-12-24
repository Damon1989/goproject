package main

import "fmt"

func f1() {
	fmt.Println("开始")
	defer func() {
		fmt.Println("aaaa")
		fmt.Println("bbbb")
	}()
	fmt.Println("结束")
}

// 调用fmt.Println(f2()) 返回0
func f2() int {
	var a int
	defer func() {
		a++
	}()
	//fmt.Println("结束")
	return a
}

// fmt.Println(f3()) 返回1
func f3() (a int) {
	defer func() {
		a++
	}()
	//fmt.Println("结束")
	return a
}

func main() {
	//fmt.Println("开始")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("结束")

	//f1()

	fmt.Println(f2())
	fmt.Println(f3())
}
