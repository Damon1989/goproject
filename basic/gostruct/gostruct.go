package gostruct

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  string
}

func structMain() {
	var p1 Person
	p1.name = "zhangsan"
	p1.age = 20
	p1.sex = "男"
	fmt.Printf("p1 name=%s age=%d sex=%s\n", p1.name, p1.age, p1.sex)
	fmt.Printf("值：%v,类型：%T \n", p1, p1)
	fmt.Printf("值：%#v,类型：%T \n", p1, p1)
}

func structMain1() {
	var p2 = new(Person)
	p2.name = "zhangsan"
	p2.age = 30
	p2.sex = "男"
	fmt.Printf("值：%v,类型：%T \n", p2, p2)
	fmt.Printf("值：%#v,类型：%T \n", p2, p2)

	fmt.Println("---------------------")
	var p3 = &Person{}
	p3.name = "zhangsan"
	p3.age = 20
	p3.sex = "女"
	fmt.Printf("值：%v,类型：%T \n", p3, p3)
	fmt.Printf("值：%#v,类型：%T \n", p3, p3)

	fmt.Println("---------------------")
	var p4 = Person{
		name: "zhangsan",
		age:  20,
		sex:  "none",
	}
	fmt.Printf("值：%v,类型：%T \n", p4, p4)
	fmt.Printf("值：%#v,类型：%T \n", p4, p4)

	fmt.Println("---------------------")
	var p5 = &Person{
		name: "zhangsan",
		age:  50,
		sex:  "none",
	}
	fmt.Printf("值：%v,类型：%T \n", p5, p5)
	fmt.Printf("值：%#v,类型：%T \n", p5, p5)

	fmt.Println("---------------------")
	var p6 = &Person{
		name: "zhangsan",
	}
	fmt.Printf("值：%v,类型：%T \n", p6, p6)
	fmt.Printf("值：%#v,类型：%T \n", p6, p6)

	fmt.Println("---------------------")
	var p7 = &Person{
		"zhangsan7",
		50,
		"none",
	}
	fmt.Printf("值：%v,类型：%T \n", p7, p7)
	fmt.Printf("值：%#v,类型：%T \n", p7, p7)
}
