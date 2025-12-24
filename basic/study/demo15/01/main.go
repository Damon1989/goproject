package main

import "fmt"

func main() {
	//var a = 10
	//fmt.Printf("a的值：%v\na的类型%T \na的地址%p", a, a, &a)

	//var a = 10
	//var p = &a
	//fmt.Printf("a的值：%v\na的类型%T \na的地址%p\n", a, a, &a)
	//fmt.Printf("p的值：%v\np的类型%T \n", p, p)

	// golang里面变量都有一个对应的内存地址，可以通过&变量名来获取
	var a = 10
	var p = &a
	fmt.Printf("a的值：%v\na的类型%T \na的地址%p\n", a, a, &a)
	fmt.Printf("p的值：%v\np的类型%T \np的地址%p\n", p, p, &p)
}
