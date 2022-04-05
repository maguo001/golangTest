package main

import (
	"fmt"
)

func main() {
	//声明一个变量接收用户选择
	key := ""
	//声明一个变量退出软件
	tuu := true
	//定义账户余额
	date := 200.0
	//定义一个变量，来判断当前有没有收入或支出
	flag := false
	//定义每次收入金额
	money := 0.0
	//定义收支说明
	spaek := ""
	//当有收支时，只需要对sum进行拼接处理
	sum := "收入\t账户余额\t收支金额\t说  明"
	//显示主菜单
	for {
		fmt.Println("\n\n------------------家庭收支记账软件------------------")
		fmt.Println("                       1 收支明细")
		fmt.Println("                       2 登记收入")
		fmt.Println("                       3 登记支出")
		fmt.Println("                       4 退出软件")
		fmt.Printf("请选择（1-4）：")
		fmt.Scanln(&key)
		fmt.Println()
		switch key {
		case "1":
			fmt.Println("-----------------------当前收支明细记录-----------------------")
			if flag {
				fmt.Println(sum)
			} else {
				fmt.Println("请来一笔收入或支出吧！！！")
			}

		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			date += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&spaek)
			//将这个收入情况拼接到sum
			sum += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", date, money, spaek)
			flag = true
		case "3":
			fmt.Println("登记支出·····")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > date {
				fmt.Println("余额不足")
				break
			}
			date -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&spaek)
			//将这个收入情况拼接到sum
			sum += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", date, money, spaek)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？ y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你输入有误！！！请重新输入 y/n")
			}
			if choice == "y" {
				tuu = false
			}
		default:
			fmt.Println("请输入正确选项！！！")
		}
		if !tuu {
			break
		}
	}
	fmt.Println("你已经退出了家庭记账软件的使用，哒哒哒")
}
