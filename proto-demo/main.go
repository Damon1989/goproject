package main

import (
	"fmt"
	"proto-demo/service"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &service.User{
		Name: "damon",
		Age:  18,
	}
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &service.User{}
	err1 := proto.Unmarshal(marshal, newUser)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(newUser.String())
}
