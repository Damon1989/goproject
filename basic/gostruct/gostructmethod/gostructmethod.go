package gostructmethod

import "fmt"

type Person struct {
	Name   string
	Age    int
	Sex    string
	Height int
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v 年龄:%v \n", p.Name, p.Age)
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func methodMain() {
	var p1 = Person{
		Name: "damon",
		Age:  18,
		Sex:  "male",
	}
	p1.PrintInfo()
	p1.SetInfo("andy", 30)
	p1.PrintInfo()

	var p2 = Person{
		Name: "olaf",
		Age:  10,
		Sex:  "male",
	}
	p2.PrintInfo()
	p1.PrintInfo()
}

type MyInt int

func (i MyInt) PrintInfo() {
	fmt.Println("我是自定义类型里面的自定义方法")
}

func myIntMain() {
	var a MyInt = 20
	a.PrintInfo()
}
