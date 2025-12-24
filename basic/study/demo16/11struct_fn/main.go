package main

import "fmt"

type User struct {
	Username string
	Password string
	Address  Address
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

	fmt.Println(u)
	fmt.Printf("%#v\n", u)
}
