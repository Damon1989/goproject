package main

import "fmt"

type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话，必须通过结构体或者自定义类型实现接口
type Phone struct {
	Name string
}

/*
*
结构体值接受者实现接口
值接受者：如果结构体中的方法是值接受者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
*/
func (p Phone) start() {
	fmt.Println(p.Name, "手机开始工作")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "手机停止工作")
}

func main() {
	var phone = Phone{
		Name: "小米",
	}
	fmt.Println(phone)

	var p1 Usber = phone
	p1.start()

	var p3 = &Phone{
		Name: "华为",
	}
	p3.start()
	var p4 Usber = p3
	p4.start()

}
