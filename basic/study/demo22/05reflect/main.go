package main

import (
	"fmt"
	"reflect"
)

// 反射获取任意变量的类型

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v)
	//v.Kind() 获取种类
	kind := v.Kind()
	switch kind {
	case reflect.Float32, reflect.Float64:
		fmt.Printf("v is float , v = %f \n", v.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("v is int , v = %d \n", v.Int())
	case reflect.String:
		fmt.Printf("v is string , v = %s \n", v.String())
	default:
		fmt.Printf("v is %v \n", kind)
	}
}
func main() {
	var a float32 = 3.14
	var b int64 = 100
	var c string = "你好golang"
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
}
