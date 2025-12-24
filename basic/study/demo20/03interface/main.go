package main

import "fmt"

type A interface{} // 空接口 表示没有任何约束  任何类型都实现了空接口

func main() {
	var a A
	var str = "你好golang"
	a = str
	fmt.Printf("值：%v 类型：%T\n", a, a)

	var num = 20
	a = num
	fmt.Printf("值：%v 类型：%T\n", a, a)

	var flag = true
	a = flag
	fmt.Printf("值：%v 类型：%T\n", a, a)
}
