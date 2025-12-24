package main

import (
	"fmt"
	"reflect"
)

// 反射获取任意变量的类型

func reflectValue(x interface{}) {
	fmt.Println(x)
	//var num = 10 + x  //invalid operation: 10 + x (mismatched types int and interface{})
	//fmt.Println(num)
	b := x.(int)
	var num = 10 + b
	fmt.Println(num)

	// 反射来实现这个功能

	//v := reflect.ValueOf(x)
	//fmt.Println(v)

	/*	var n = v + 12
		fmt.Println(n)*/
	// 反射获取变量的原始值
	v := reflect.ValueOf(x)
	m := v.Int() + 12
	fmt.Println(m)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int() 从反射中获取整型的原始值
		fmt.Printf("type is int64, value is %d \n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f \n", v.Float())
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f \n", v.Float())
	}
}
func main() {
	var a = 13
	reflectValue(a)
}
