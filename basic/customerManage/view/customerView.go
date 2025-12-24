package main

import (
	"customerManagement/model"
	"customerManagement/service"
	"fmt"
)

type CustomerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (this *CustomerView) list() {
	customers := this.customerService.List()
	fmt.Println("----------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------客户列表完成----------")
}

func (this *CustomerView) add() {
	fmt.Println("添加客户...")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("添加客户成功")
	} else {
		fmt.Println("添加客户失败")
	}
}

func (this *CustomerView) update() {}

func (this *CustomerView) delete() {
	fmt.Println("删除客户...")
	fmt.Println("请输入要删除的客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(y/n):")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入y/n")
	}
	if choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("删除客户成功")
		} else {
			fmt.Println("删除客户失败，客户编号不存在")
		}
	}
}

func (this *CustomerView) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *CustomerView) mainMenu() {
	for {
		println("----------客户信息管理软件----------")
		println("            1 添 加 客 户")
		println("            2 修 改 客 户")
		println("            3 删 除 客 户")
		println("            4 客 户 列 表")
		println("            5 退 出")
		println("请选择(1-5):")

		// 等待用户输入
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添 加 客 户...")
			this.add()
		case "2":
			fmt.Println("修 改 客 户...")
		case "3":
			fmt.Println("删 除 客 户...")
		case "4":
			fmt.Println("客 户 列 表...")
			this.list()
		case "5":
			// 让用户确认是否退出
			fmt.Println("你确定要退出吗？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入y/n")
			}
			if choice == "y" {
				this.loop = false
			}
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户信息管理软件的使用...")
}

func main() {
	customerView := CustomerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
