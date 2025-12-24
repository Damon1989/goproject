package gojson

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func jsonMain() {
	var s1 = Student{
		ID:     12,
		Gender: "male",
		Name:   "李四",
		Sno:    "s0010",
	}
	fmt.Printf("%#v\n", s1)

	jsonByte, _ := json.Marshal(s1)
	jsonString := string(jsonByte)
	fmt.Println(jsonString)

	var str = `{"ID":12,"Gender":"male","Name":"李四","Sno":"s0010"}`
	var stu1 Student
	err := json.Unmarshal([]byte(str), &stu1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu1)
	fmt.Printf("%#v\n", stu1)
}
