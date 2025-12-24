package main

import "fmt"

type User struct {
	Username string
	Password string
	Address
}

type Address struct {
	Name  string
	Phone string
	City  string
}

func main() {
	var u User
	u.Username = "admin"
	u.Password = "123456"
	u.Address.Name = "张三"
	u.Address.Phone = "12345678901"
	u.Address.City = "北京"

	u.City = "上海"
	fmt.Println(u)
	fmt.Printf("%#v\n", u) //main.User{Username:"admin", Password:"123456", Address:main.Address{Name:"张三", Phone:"12345678901", City:"上海"}}

	fmt.Println(u.Address.Phone)
	fmt.Println(u.Phone)
}
