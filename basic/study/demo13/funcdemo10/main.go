package main

import "fmt"

func fn1(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}

func fn2(n int) int {
	if n > 1 {
		return n + fn2(n-1)
	}
	return 1
}

// 递归实现5的阶乘
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	fn1(10)

	fmt.Println("*****************")
	fmt.Println(fn2(100))

	fmt.Println("*****************")
	fmt.Println(factorial(5))
}
