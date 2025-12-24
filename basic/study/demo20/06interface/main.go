package main

import "fmt"

type Usber interface {
	Start()
	Stop()
}

// 如果接口里面有方法的话，必须通过结构体或者自定义类型实现接口
type Phone struct {
	Name string
}

// 手机要实现Usber接口的方法

func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

func (c Camera) run() {
	fmt.Println("相机running")
}

type Computer struct {
}

func (c Computer) work(usb Usber) {
	if usb, ok := usb.(Camera); ok {
		usb.run()
	}
	if usb, ok := usb.(Phone); ok {
		usb.Start()
	}

}

func main() {
	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	var camera = Camera{}
	//var camera=Camera{}
	computer.work(phone)
	computer.work(camera)
}
