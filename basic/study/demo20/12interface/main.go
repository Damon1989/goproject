package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// Golang中空接口和类型断言使用细节
func main() {
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 18
	userinfo["sex"] = true
	userinfo["hobby"] = []string{"篮球", "足球", "乒乓球"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["hobby"])
	fmt.Println(userinfo["hobby"].([]string)[0])

	var address = Address{
		Name:  "张三",
		Phone: 12345678901,
	}
	fmt.Println(address.Name)
	userinfo["address"] = address

	fmt.Println(userinfo["address"])

	var name = userinfo["address"].(Address).Name
	fmt.Println(name)
}
