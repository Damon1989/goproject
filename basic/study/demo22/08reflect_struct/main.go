package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法...")
}

func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的参数不是一个指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体指针类型")
		return
	}

	// 修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("小红")

	age := v.Elem().FieldByName("Age")
	age.SetInt(20)
}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   18,
		Score: 99,
	}
	fmt.Println(stu1)
	reflectChangeStruct(&stu1)
	fmt.Println(stu1)

	/*	var a = 12
		reflectChangeStruct(&a)*/

}
