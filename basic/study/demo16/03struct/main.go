package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	var p2 = new(Person)
	p2.Name = "李四"
	p2.Age = 20
	p2.Sex = "男"
	(*p2).Name = "王五"
	fmt.Printf("值：%#v 类型：%T 地址：%p\n", p2, p2, &p2) //值：&main.Person{Name:"李四", Age:20, Sex:"男"} 类型：*main.Person 地址：0xc00006c028

	var p3 = &Person{}
	p3.Name = "赵六"
	p3.Age = 22
	p3.Sex = "男"
	(*p3).Name = "孙七"
	fmt.Printf("值：%#v 类型：%T 地址：%p\n", p3, p3, &p3) //值：&main.Person{Name:"孙七", Age:22, Sex:"男"} 类型：*main.Person 地址：0xc00006c038

	var p4 = Person{
		Name: "周八",
		Age:  24,
		Sex:  "男",
	}
	fmt.Println("------------------------")
	fmt.Printf("值：%#v 类型：%T 地址：%p\n", p4, p4, &p4) //值：main.Person{Name:"周八", Age:24, Sex:"男"} 类型：main.Person 地址：0xc0000261b0

	fmt.Println("------------------------")

	var p5 = &Person{
		Name: "周八",
		Age:  24,
		Sex:  "男",
	}
	fmt.Printf("值：%#v 类型：%T 地址：%p\n", p5, p5, &p5)

	fmt.Println("------------------------")

	var p6 = Person{
		Name: "周八11",
	}
	fmt.Printf("值：%#v 类型：%T 地址：%p\n", p6, p6, &p6)

	fmt.Println("------------------------")

	var p7 = &Person{
		"周八22",
		26,
		"男",
	}
	fmt.Printf("值：%#v 类型：%T 地址：%p\n", p7, p7, &p7)

}
