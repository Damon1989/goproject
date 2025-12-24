package main

import "fmt"

var a = "全局变量"

func run() {
	b := "局部变量"
	fmt.Println(b)
	fmt.Println("run方法a=", a)
	fmt.Println("run方法b=", b)
}
func main() {
	fmt.Println("main方法a=", a)
	run()
}
