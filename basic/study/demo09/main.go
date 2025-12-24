package main

import "fmt"

func main() {
	//1、数组的长度是类型的一部分
	//var arr1 [3]int
	//var arr2 [4]int
	//var strArr [3]string
	//
	//fmt.Printf("arr1:%T;arr2:%T;strArr:%T", arr1, arr2, strArr)

	////	2.数组的初始化 第一种方法
	//var arr1 [3]int
	//fmt.Println(arr1)

	var slice []int

	slice = make([]int, 5, 10)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5

	fmt.Println(slice)

	fmt.Println("slice的长度是：", len(slice))
	fmt.Println("slice的容量是：", cap(slice))
}
