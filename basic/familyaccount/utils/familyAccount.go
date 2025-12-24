package utils

import "fmt"

type FamilyAccount struct {
	// 必须声明的字段

	// 声明一个字段，保存接收用户输入的选项
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	details string
}

// 编写要给工厂模式的构造方法，返回一个 *FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说  明\n",
	}
}

// 给FamilyAccount绑定方法
// 显示明细
func (fa *FamilyAccount) ShowDetails() {
	if fa.flag {
		println("----------当前收支明细记录----------")
		println(fa.details)
	} else {
		println("----------当前没有收支明细记录----------")
	}
}

// 登记收入
func (fa *FamilyAccount) Income() {
	println("本次收入金额：")
	fmt.Scanln(&fa.money)
	fa.balance += fa.money
	println("本次收入说明：")
	fmt.Scanln(&fa.note)
	// 将收入情况，拼接到details变量
	fa.details += fmt.Sprintf("收入\t%v\t%v\t%v\n", fa.balance, fa.money, fa.note)
	fa.flag = true // 修改标记
}

// 登记支出
func (fa *FamilyAccount) Pay() {
	println("本次支出金额：")
	fmt.Scanln(&fa.money)
	// 这里需要做一个必要的判断
	if fa.money > fa.balance {
		println("支出金额大于账户余额，支付失败")
		return // 直接返回，不在继续执行
	}
	fa.balance -= fa.money
	println("本次支出说明：")
	fmt.Scanln(&fa.note)
	// 将支出情况，拼接到details变量
	fa.details += fmt.Sprintf("支出\t%v\t%v\t%v\n", fa.balance, fa.money, fa.note)
	fa.flag = true // 修改标记
}

// 退出软件
func (fa *FamilyAccount) Exit() {
	println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		println("你的输入有误，请重新输入y/n")
	}
	if choice == "y" {
		fa.loop = false
	}
}

// 给FamilyAccount绑定方法
// 显示主菜单
func (fa *FamilyAccount) MainMenu() {
	for fa.loop {
		// 显示菜单
		println("----------家庭收支记账软件----------")
		println("            1 收支明细")
		println("            2 登记收入")
		println("            3 登记支出")
		println("            4 退出软件")
		println("请选择(1-4):")

		// 等待用户输入
		fmt.Scanln(&fa.key)
		switch fa.key {
		case "1":
			fa.ShowDetails()
		case "2":
			fa.Income()
		case "3":
			fa.Pay()
		case "4":
			fa.Exit()
		default:
			println("你的输入有误，请重新输入")
		}
	}

	println("你退出了家庭收支记账软件的使用...")
}
