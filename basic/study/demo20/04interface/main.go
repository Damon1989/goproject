package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("值：%v 类型：%T\n", a, a)
}

// golang中空接口也可以直接当做类型使用，可以表示任意类型
func main() {
	//var a interface{}
	//
	//a = 20
	//fmt.Printf("值：%v 类型：%T\n", a, a)
	//
	//a = "你好golang"
	//fmt.Printf("值：%v 类型：%T\n", a, a)
	//
	//a = true
	//fmt.Printf("值：%v 类型：%T\n", a, a)

	/*	show(20)
		show("你好golang")
		show(true)
		slice := []interface{}{20, "你好golang", true}
		for _, v := range slice {
			fmt.Printf("值：%v 类型：%T\n", v, v)
		}
		show(slice)*/

	var m1 = make(map[string]interface{})
	m1["name"] = "小明"
	m1["age"] = 18
	m1["height"] = 1.88
	m1["married"] = false
	for k, v := range m1 {
		fmt.Printf("key=%v value=%v\n", k, v)
	}

	fmt.Println("------------------------------")

	var s1 = []interface{}{"小明", 18, 1.88, false}
	for _, v := range s1 {
		fmt.Printf("value=%v\n", v)
	}
}
