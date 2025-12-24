package main

import "fmt"

type calc func(int, int) int //表示顶一个calc类型

type myInt int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func test() {
	fmt.Println("test ...")
}

func main() {
	var c calc
	c = sub
	fmt.Printf("c的类型：%T\n", c)

	fmt.Println(c(10, 5))

	var f = sub
	fmt.Printf("f的类型：%T\n", f)
	fmt.Println(f(15, 5))

	var a int = 10
	var b myInt = 20
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("a的类型：%T\n", a)
	fmt.Printf("b的类型：%T\n", b)
	fmt.Println(a + int(b))
}
