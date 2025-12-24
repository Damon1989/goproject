package main

import "fmt"

func main() {
	// 声明一个变量，保存接收用户输入的选项
	key := ""
	// 声明一个变量，控制是否退出for循环
	loop := true

	// 定义账户的余额[]
	balance := 10000.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 定义个变量,记录是否有收支的行为
	flag := false
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对details进行拼接处理即可
	details := "收支\t账户金额\t收支金额\t说    明"

	for loop {
		// 显示菜单
		println("----------家庭收支记账软件----------")
		println("            1 收支明细")
		println("            2 登记收入")
		println("            3 登记支出")
		println("            4 退出软件")
		println("请选择(1-4):")

		// 等待用户输入
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("显示收支明细...")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细... 来一笔吧！")
			}
			fmt.Println(details)
		case "2":
			fmt.Println("登记收入...")
			fmt.Scanln(&money)
			balance += money //修改账户金额
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			// 将这个收入情况，拼接到details变量
			//  收入 11000  1000  有人发红包
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("登记支出...")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			// 这里需要做一个必要的判断
			if money > balance {
				fmt.Println("余额的金额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	fmt.Println("你退出了家庭收支记账软件的使用...")
}
