package main

import (
	"fmt"
	"olaf/calc"
	"olaf/tools"
)

func init() {
	fmt.Println("main init")
}

func main() {
	sum := calc.Add(100, 200)
	sub := calc.Sub(200, 100)
	fmt.Println(sum, sub)
	fmt.Println(calc.Age)

	mul := tools.Mul(2, 3)
	div := tools.Div(6, 3)
	fmt.Println(mul, div)

	tools.PrintInfo()
	tools.SortIntAsc()
}
