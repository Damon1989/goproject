package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 Person //实例化Person结构体
	p1.name = "张三"
	p1.sex = "男"
	p1.age = 18
	fmt.Printf("值：%v 类型：%T 地址：%p\n", p1, p1, &p1)
	fmt.Printf("值：%#v 类型：%T 地址：%p\n", p1, p1, &p1) //值：main.Person{name:"张三", age:18, sex:"男"} 类型：main.Person 地址：0xc0000880c0

}
