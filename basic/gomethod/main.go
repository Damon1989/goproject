package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) say() {
	println("My name is", p.name, "and I am", p.age, "years old.")
}
func (p Person) setName(name string) {
	p.name = name
}
func (p *Person) setAge(age int) {
	p.age = age
}

func say() {
	println("I am a function.")
}

func test(p Person) {
	fmt.Println("my name is ", p.name, " my age is ", p.age, "!!!")
}

func main1() {
	var p = Person{"Tom", 18}
	//p.say()
	//say()
	//fmt.Printf("%T\n", p.say)
	//fmt.Printf("%T\n", say)

	//println("-------------------------------")
	//
	//var fn func()
	//fn = p.say
	//fn()
	//
	//fn = say
	//fn()

	fmt.Println("-------------------------------")
	test(p)
}

func main2() {
	per := Person{"Tom", 18}
	fmt.Println(per)
	per.setName("Jerry")
	fmt.Println(per)
	p := &per
	(*p).setAge(33)
	fmt.Println(per)
}

func main() {
	per := Person{"Tom", 18}
	// 方式一：先拿到指针，再调用方法
	p := &per
	(*p).setAge(33)
	fmt.Println(per)

	// 方式二：直接调用方法,底层会自动获取变量地址传递给接受者
	per.setAge(44)
	fmt.Println(per)
}
