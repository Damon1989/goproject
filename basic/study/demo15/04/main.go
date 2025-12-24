package main

import "fmt"

func main() {
	//var userinfo = make(map[string]string)
	//userinfo["username"] = "张三"
	//fmt.Println(userinfo)
	//
	//var a = make([]int, 4, 4)
	//a[0] = 1
	//fmt.Println(a)

	var a *int = new(int)
	*a = 100
	fmt.Println(*a)
}
