package main

import (
	"fmt"
	"os"

	"github.com/ZhijiunY/MUCS/client/process"
)

var userId int
var userPwd string
var userName string

func main() {

	var key int

	for {
		fmt.Println("----------  歡迎登入多人聊天系統  ----------")
		fmt.Println("\t\t\t 1 登入聊天室")
		fmt.Println("\t\t\t 2 註冊用戶")
		fmt.Println("\t\t\t 3 退出系統")
		fmt.Println("\t\t\t 請選擇（1~3）:")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登入聊天室")
			fmt.Println("請輸入用戶的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("請輸入用戶的密碼")
			fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("註冊用戶")
			fmt.Println("請輸入用戶id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("請輸入用戶密碼：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("請輸入用戶名字（nickname）:")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系統")
			os.Exit(0)
		default:
			fmt.Println("你的輸入有誤，請重新輸入")
		}
	}

}
