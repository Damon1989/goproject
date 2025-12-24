package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(100)
	}
}

func reflectSetValue(x interface{}) {
	//*x = 120 invalid operation: cannot indirect x (variable of type interface{})

	//v, _ := x.(*int)
	//*v = 120  invalid memory address or nil pointer dereference

	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())        //ptr
	fmt.Println(v.Elem().Kind()) //int64
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(120)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("hello")
	}
}

func main() {
	var a int64 = 100
	//reflectSetValue1(&a)
	reflectSetValue(&a)

	fmt.Println(a)

	fmt.Println("*****************")
	var b string = "你好golang"
	reflectSetValue(&b)
	fmt.Println(b)
}
