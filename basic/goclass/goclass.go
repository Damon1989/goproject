package goclass

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func ClassDemo01() {
	r1 := Rect{Point{1, 1}, Point{2, 2}}
	fmt.Printf("r1.leftUp.x 地址=%p\n", &r1.leftUp.x)
	fmt.Printf("r1.leftUp.y 地址=%p\n", &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x 地址=%p\n", &r1.rightDown.x)
	fmt.Printf("r1.rightDown.y 地址=%p\n", &r1.rightDown.y)

	fmt.Println("=========================================")

	r2 := Rect2{&Point{1, 1}, &Point{2, 2}}
	fmt.Printf("r2.leftUp.x 地址=%p\n", &r2.leftUp.x)
	fmt.Printf("r2.leftUp.y 地址=%p\n", &r2.leftUp.y)
	fmt.Printf("r2.rightDown.x 地址=%p\n", &r2.rightDown.x)
	fmt.Printf("r2.rightDown.y 地址=%p\n", &r2.rightDown.y)

	fmt.Printf("r1.leftUp 地址=%p\n", &r1.leftUp)
	fmt.Printf("r1.rightDown 地址=%p\n", &r1.rightDown)
}
