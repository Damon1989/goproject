package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 自定义一个方法类型
type calcType func(int, int) int

func calc(x, y int, op calcType) int {
	return op(x, y)
}

func main() {
	sum := calc(10, 5, add)
	fmt.Println(sum)

	s := calc(10, 5, sub)
	fmt.Println(s)

	j := calc(3, 4, func(x int, y int) int {
		return x * y
	})
	fmt.Println(j)
}
