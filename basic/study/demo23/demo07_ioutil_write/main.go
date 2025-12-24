package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello golang"
	err := ioutil.WriteFile("./log.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
