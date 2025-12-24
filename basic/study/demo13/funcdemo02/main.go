package main

import "fmt"

func calc(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc1(x int, y int) (sum int, sub int) {
	fmt.Println(sum, sub)
	sum = x + y
	sub = x - y
	fmt.Println(sum, "------", sub)
	return
}

func calc2(x, y int) (sum, sub int) {
	fmt.Println(sum, sub)
	sum = x + y
	sub = x - y
	fmt.Println(sum, "------", sub)
	return
}

func main() {
	//sum, sub := calc(3, 2)
	//fmt.Println(sum, sub)
	//
	//sum1, sub1 := calc1(10, 2)
	//fmt.Println(sum1, sub1)
	//
	//sum2, sub2 := calc2(10, 2)
	//fmt.Println(sum2, sub2)
	//

	sum2, _ := calc2(10, 2)
	fmt.Println(sum2)

}
