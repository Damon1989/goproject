package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteStr, err := ioutil.ReadFile("./mainold.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}
