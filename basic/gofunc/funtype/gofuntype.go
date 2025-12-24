package funtype

import "fmt"

type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func typeMain() {
	var c calc
	c = add
	fmt.Printf("c的类型：%T\n", c)

	f := sub
	fmt.Printf("f的类型：%T\n", f)
	fmt.Println(f(10, 2))
}
