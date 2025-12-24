package main

import (
	"fmt"
	"reflect"
)

// 反射获取任意变量的类型

type MyInt int
type Person struct {
	Name string
	Age  int
}

func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	//v.Name()  // 类型名称  种类（Kind）就是指底层的类型
	//v.Kind()  // 类型种类
	fmt.Printf("类型：%v 类型名称：%v 类型种类：%v\n", v, v.Name(), v.Kind())
}
func main() {
	a := 10         //int
	b := 23.4       //float64
	c := true       //bool
	d := "你好golang" //string
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e MyInt = 34
	var f = Person{
		Name: "张三",
		Age:  18,
	}
	reflectFn(e) //类型：main.MyInt 类型名称：MyInt 类型种类：int
	reflectFn(f) //类型：main.Person 类型名称：Person 类型种类：struct

	var h = 25
	reflectFn(&h) //类型：*int 类型名称： 类型种类：ptr

	var i = [3]int{1, 2, 3}   //类型：[3]int 类型名称： 类型种类：array
	var j = []int{11, 22, 33} //类型：[]int 类型名称： 类型种类：slice
	reflectFn(i)
	reflectFn(j)
}
