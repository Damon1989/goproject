package main

import "fmt"

func main() {
	//var a = new(int) // a 是一个指针变量 类型 *int  值是0
	//
	//fmt.Printf("a的值：%v\n 类型：%T\n  指针变量对应的值：%v", a, a, *a)

	/**
		错误的写法
		var a *int
	*a = 100
	fmt.Println(*a)
	*/

	/*	var a *int
		a = new(int)
		*a = 100
		fmt.Println(*a)*/

	var f = new(bool)
	fmt.Println(*f)
}
