package main

import "fmt"

// 定义一个Animaler的接口，Animal中定义两个方法，分别是SetName和GetName,分别让Dog和Cat实现Animal接口
type Animaler1 interface {
	SetName(name string)
}

type Animaler2 interface {
	GetName() string
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
	var d1 Animaler1 = d //表示让Dog实现Animaler1接口
	d1.SetName("旺财2")

	var d2 Animaler2 = d //表示让Dog实现Animaler2接口
	fmt.Println(d2.GetName())

}
