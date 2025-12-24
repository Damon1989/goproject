package main

import "fmt"

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

func subFn(x int, y int) int {
	fmt.Println("x:", x, "y:", y)
	sub := x - y
	return sub
}

// 参数简写
func subFn1(x, y int) int {
	fmt.Println("x:", x, "y:", y)
	sub := x - y
	return sub
}

func sumFn1(x ...int) int {
	fmt.Printf("%v -- %T\n", x, x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}

func main() {

	sum1 := sumFn(12, 3)
	fmt.Println(sum1) //15

	sum2 := sumFn(15, 5)
	fmt.Println(sum2) //20

	a := 20
	b := 2
	sub1 := subFn(a, b)
	fmt.Println(sub1) //18

	sub2 := subFn1(a, b)
	fmt.Println(sub2) //18

	fn1 := sumFn1(1, 2, 3, 4)
	fmt.Println(fn1) //10
}
