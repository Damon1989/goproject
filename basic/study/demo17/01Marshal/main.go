package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	name   string //私有数学不能被json包访问
	Sno    string
}

func main() {
	var s1 = Student{
		ID:     1,
		Gender: "男",
		name:   "张三",
		Sno:    "S0001",
	}
	fmt.Printf("s1:%#v\n", s1)

	jsonByte, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json jsonByte failed")
		return
	}
	fmt.Printf("s1:%#v\n", string(jsonByte))
	fmt.Println(string(jsonByte)) // {"ID":1,"Gender":"男","Name":"张三","Sno":"S0001"}
}
