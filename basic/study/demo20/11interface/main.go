package main

import "fmt"

// 定义一个Animaler的接口，Animal中定义两个方法，分别是SetName和GetName,分别让Dog和Cat实现Animal接口
type Ainterface interface {
	SetName(name string)
}

type Binterface interface {
	GetName() string
}

type Animaler interface { //接口的嵌套
	Ainterface
	Binterface
}

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}
func (d Dog) GetName() string {
	return d.name
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

}
