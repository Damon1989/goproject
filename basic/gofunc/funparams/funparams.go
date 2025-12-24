package funparams

import (
	"fmt"
	"log"
)

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

func methodMain() {
	sum := calc(10, 5, add)
	log.Println(sum)
	s := calc(10, 5, sub)
	log.Println(s)

	j := calc(3, 4, func(x int, y int) int {
		return x + y
	})
	fmt.Println(j)
}

func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x int, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func doMain() {
	var a = do("+")
	sum := a(1, 2)
	log.Println(sum)

	b := do("*")
	result := b(3, 4)
	log.Println(result)
}

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

func fn3(n int) int {
	if n > 1 {
		return n * fn3(n-1)
	}

	return 1

}
