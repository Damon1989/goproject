package main

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

func main() {
	var student = Student{
		ID:     1,
		Gender: "男",
		Name:   "张三",
		Sno:    "20190001",
	}
	fmt.Printf("student:%#v\n", student)
	jsonByte, _ := json.Marshal(student)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

}
