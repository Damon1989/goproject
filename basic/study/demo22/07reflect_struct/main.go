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

// 打印字段
func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	field0 := t.Field(0)
	fmt.Println(field0)         //{Name  string json:"name" 0 [0] false}
	fmt.Printf("%v\n", field0)  //{Name  string json:"name" 0 [0] false}
	fmt.Printf("%#v\n", field0) //reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x603da0), Tag:"json:\"name\"", Offset:0x0, Index:[]int{0}, Anonymous:false}

	fmt.Println("-----------------")

	fmt.Println("自动名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段标签：", field0.Tag)
	fmt.Println("字段标签：", field0.Tag.Get("json"))

	fmt.Println("**************************************")

	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("自动名称：", field1.Name)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段标签：", field1.Tag)
		fmt.Println("字段标签：", field1.Tag.Get("json"))
	}

	fmt.Println("////////////**************************************")

	fieldCount := t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")
	/*	for i := 0; i < fieldCount; i++ {
		field := t.Field(i)
		fmt.Println("自动名称：", field.Name)
		fmt.Println("字段类型：", field.Type)
		fmt.Println("字段标签：", field.Tag)
		fmt.Println("字段标签：", field.Tag.Get("json"))
	}*/

	fmt.Println("0000000000000000000000000000")
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println(v.FieldByName("Score"))

	fmt.Println("000000000000000000000000000011111")
	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称：%v  属性值：%v 属性类型：%v 属性Tag: %v \n", t.Field(i).Name,
			v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}
}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	method0 := t.Method(0)    // 和结构体方法的顺序没有关系，和结构体方法的AsCII有关系
	fmt.Println(method0.Name) //GetInfo
	fmt.Println(method0.Type) //func(main.Student) string

	fmt.Println("----------------------------")
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name) //Print
		fmt.Println(method1.Type) //func(main.Student)
	}

	v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil)
	info1 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info1)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("小红"))
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(100))
	v.MethodByName("SetInfo").Call(params) //执行方法传入参数

	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	fmt.Println("方法数量", t.NumMethod())
}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   18,
		Score: 99,
	}
	//PrintStructField(stu1)
	PrintStructFn(&stu1)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(stu1)
}
