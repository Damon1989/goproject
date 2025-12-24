package main

import "fmt"

// 父亲结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Println(a.Name, "is running")
}

// 子结构体
type Dog struct {
	*Animal // 结构体嵌套 继承
	Age     int
}

func (d Dog) wang() {
	fmt.Println(d.Name, "is wang")
}

func main() {
	var d = Dog{
		Animal: &Animal{
			Name: "旺财",
		},
		Age: 3,
	}
	d.run()
	d.wang()
}
