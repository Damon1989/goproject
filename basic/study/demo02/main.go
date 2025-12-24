package main

import "fmt"

func main() {
	//fmt.Println("Hello, World!")
	//fmt.Print("Hello, World!!")
	//fmt.Printf("Hello, World!!!\n")

	/*
		注释
	*/

	// 1.Print Println 区别  ctrl+/
	//fmt.Println("Hello, World!")
	//fmt.Println("Hello, World!")
	//fmt.Print("Hello, World!!")
	//fmt.Print("Hello, World!!")

	//// 2.变量
	//var a = "aaa"
	//fmt.Printf("a=%s\n", a)

	//var a int = 10
	//var b int = 3
	//var c int = 5
	////fmt.Println("a=", a, "b=", b, "c=", c)
	//fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	//fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

	a := 10
	//b := 3
	//c := 5
	//fmt.Println("a=", a, "b=", b, "c=", c)
	//fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

	//使用Printf打印一个变量的类型
	fmt.Printf("a=%v a的类型是%T\n", a, a)
}
