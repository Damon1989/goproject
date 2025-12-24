package main

import "fmt"

func main() {
	a := 10

	p := &a // p 指针变量  类型 *int

	//*p  // *p 代表取出指针变量p所指向的内存地址的值

	fmt.Println("a的值：", a)
	fmt.Println("a的地址：", &a)
	fmt.Println("p的值：", p)
	fmt.Println("p的地址：", &p)
	fmt.Println("*p的值：", *p)

	*p = 100
	fmt.Println("----------------")
	fmt.Println("a的值：", a)
	fmt.Println("*p的值：", *p)
}
