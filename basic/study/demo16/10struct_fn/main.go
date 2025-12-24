package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Hobby []string
	map1  map[string]string
}

func main() {
	var p Person
	p.Name = "张三"
	p.Age = 18
	//p.Hobby = []string{"篮球", "足球", "乒乓球"}
	p.Hobby = make([]string, 3, 6)
	p.Hobby[0] = "篮球"
	p.Hobby[1] = "足球"
	p.Hobby[2] = "乒乓球"
	p.map1 = make(map[string]string)
	p.map1["name"] = "张三"
	p.map1["age"] = "18"
	fmt.Println(p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%#v\n", p.Hobby)
}
