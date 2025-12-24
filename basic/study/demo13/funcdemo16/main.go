package main

import "fmt"

func fn1(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("恢复异常：", err)
		}
	}()
	return a / b

	panic("抛出一个异常")
}

func main() {
	fmt.Println(fn1(10, 0))
	fmt.Println("结束")
	fmt.Println(fn1(10, 2))
}
