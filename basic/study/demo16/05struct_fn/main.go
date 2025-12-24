package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Sex    string
	Height int
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s 年龄：%d 性别：%s 身高：%d\n", p.Name, p.Age, p.Sex, p.Height)
}

func main() {
	var p1 = Person{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}

	p1.PrintInfo()

	var p2 = Person{
		Name:   "李四",
		Age:    20,
		Sex:    "男",
		Height: 180,
	}

	p2.PrintInfo()
}
