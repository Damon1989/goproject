package main

import "fmt"

type usber interface {
	start()
	stop()
}

type Consumer struct {
	name  string
	model string
}

func (c Consumer) start() {
	println("启动电脑")
}

func (c Consumer) stop() {
	println("关闭电脑")
}

type Phone struct {
	name  string
	model string
}

func (p Phone) start() {
	println("启动手机")
}

func (p Phone) stop() {
	println("关闭手机")
}

func working(usb usber) {
	usb.start()
	usb.stop()
}

func main1() {
	cm := Consumer{"联想", "Y7000"}
	working(cm)

	ph := Phone{"华为", "P30"}
	working(ph)
}

func main2() {
	var i interface{}
	i = 123
	fmt.Println(i)
	i = 3.14
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)
	i = Consumer{"联想", "Y7000"}
	fmt.Println(i)
	i = Phone{"华为", "P30"}
	fmt.Println(i)
}

type studier interface {
	read()
}

type Person struct {
	name string
	age  int
}

func (p Person) read() {
	fmt.Println(p.name, "在读书")
}

func main() {
	var s studier
	s = Person{"Tom", 18}
	s.read()
	if p, ok := s.(Person); ok {
		p.name = "Jerry"
		fmt.Println(p)
	}
	switch p := s.(type) {
	case Person:
		p.name = "Jerry1"
		fmt.Println(p)
	default:
		fmt.Println("类型不匹配")
	}
}
