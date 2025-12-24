package main

import "fmt"

// 自定义类型
type myInt int
type myFn func(int, int) int

// 类型别名
type myFloat = float64

func main() {
	var a myInt = 10
	fmt.Printf("a的值：%v\na的类型%T \na的地址%p\n", a, a, &a)

	var b myFloat = 10.1
	fmt.Printf("b的值：%v\nb的类型%T \nb的地址%p\n", b, b, &b)
}
