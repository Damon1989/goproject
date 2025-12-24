package main

import "fmt"

// 定义一个方法，可以传入任意数据类型，然后根据不同的类型实现不同的功能
// x.(type) 只能在switch中使用
func MyPrint(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("a是string类型")
	case int:
		fmt.Println("a是int类型")
	default:
		fmt.Println("a是未知类型")
	}
}

func main() {
	var a interface{}

	//a = "你好golang"
	a = 11
	v, ok := a.(string)
	if ok {
		fmt.Println("a是string类型,值是：", v)
	} else {
		fmt.Println("a不是string类型")
	}

	MyPrint("1")
	MyPrint(11)
	MyPrint(true)
}
