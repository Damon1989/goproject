package main

import (
	"fmt"
	"unsafe"
)

func main1() {
	var p1 *int
	var p2 *float64
	var p3 *bool
	fmt.Println(unsafe.Sizeof(p1)) // 8
	fmt.Println(unsafe.Sizeof(p2)) // 8
	fmt.Println(unsafe.Sizeof(p3)) // 8
}

func main2() {
	num := 666
	var p *int = &num
	fmt.Printf("%p\n", &num)
	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", *p)

	*p = 888
	fmt.Println(num)
	fmt.Println(*p)
}

func main3() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])

}

func main4() {
	var arr [3]int = [3]int{1, 2, 3}
	var p1 *[3]int = &arr
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", p1)
	//p1=arr  // cannot use arr (type [3]int) as type *[3]int in assignment
	p1[1] = 666
	fmt.Println(arr[1])

	fmt.Println("---------------------------------")
	var p2 *[3]int = &arr
	fmt.Printf("%T\n", &arr)
	fmt.Printf("%T\n", p2)
	p2[1] = 888
	fmt.Println(arr[1])

}

func main() {
	var sce []int = []int{1, 2, 3}
	fmt.Printf("sce = %p\n", sce)
	fmt.Println(sce)
}
