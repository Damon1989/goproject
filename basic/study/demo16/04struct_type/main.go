package main

import "fmt"

type Person struct {
	Name string
	Age  string
	Sex  string
}

func main() {
	var p1 = Person{
		Name: "张三",
		Age:  "18",
		Sex:  "男",
	}

	p2 := p1
	p2.Name = "李四"
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

}
