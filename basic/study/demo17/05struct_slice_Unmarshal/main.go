package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title    string
	Students []Student
}

func main() {
	// 创建一个班级变量
	c := Class{
		Title: "101",
		Students: []Student{
			{1, "男", "张三"},
			{2, "女", "李四"},
			{3, "男", "王五"},
		},
	}
	fmt.Printf("%#v\n", c)

	strByte, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("%v\n", string(strByte)) //{"Title":"101","Students":[{"ID":1,"Gender":"男","Name":"张三"},{"ID":2,"Gender":"女","Name":"李四"},{"ID":3,"Gender":"男","Name":"王五"}]}

	var str = `{"Title":"101","Students":[{"ID":1,"Gender":"男","Name":"张三"},{"ID":2,"Gender":"女","Name":"李四"},{"ID":3,"Gender":"男","Name":"王五"}]}`
	var d = Class{}
	json.Unmarshal([]byte(str), &d)
	fmt.Println("-----------------")
	fmt.Printf("%#v\n", d)
	fmt.Printf("%v\n", d)

}
