package main

import "fmt"

// 定义一个Animaler的接口，Animal中定义两个方法，分别是SetName和GetName,分别让Dog和Cat实现Animal接口
type Animaler interface {
	SetName(name string)
	GetName() string
}
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}
func (d Dog) GetName() string {
	return d.name
}
func (c *Cat) SetName(name string) {
	c.name = name
}
func (c Cat) GetName() string {
	return c.name
}
func main() {
	//Dog实现Animaler接口
	d := &Dog{
		name: "旺财",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("旺财2")
	fmt.Println(d1.GetName())

	c := &Cat{
		name: "小花",
	}
	var c1 Animaler = c
	fmt.Println(c1.GetName())
	c1.SetName("小花2")
	fmt.Println(c1.GetName())
}
