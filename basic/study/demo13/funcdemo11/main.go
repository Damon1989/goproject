package main

import "fmt"

/**

全局变量特点：
1.常驻内存
2.污染全局

局部变量特点：
1.不常驻内存
2.不污染全局

闭包：
1.可以让一个变量常驻内存
2.可以让一个变量不污染全局
*/

//var a = 12
//
//func test() {
//	var a = 3
//	fmt.Println(a)
//}

func adder1() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func() int {
	var i = 10
	return func() int {
		i++
		return i + 1
	}
}

func main() {
	//test()
	//fmt.Println(a)
	var fn1 = adder1()
	fmt.Println(fn1())
	fmt.Println(fn1())
	fmt.Println(fn1())

	var fn2 = adder2()
	fmt.Println(fn2())
	fmt.Println(fn2())
	fmt.Println(fn2())
}
