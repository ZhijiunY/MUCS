package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ZhijiunY/MUCS/server/model"
)

// 處理和客戶端的通訊
func Process(conn net.Conn) {

	// 這裡需要延時關閉conn
	defer conn.Close()

	// 這裡調用總控，創建一個
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客戶端和服務器通訊協成錯誤=err", err)
		return
	}
}

// 這裡我們編寫一個函數，完成對 UserDao 的初始化任務
func initUserDao() {
	// 這裡的 pool 本身就是一個全局變量
	// 這裡需要注意一個初始化順序問題
	// initPool, 在 initUserDao

	model.MyUserDao = model.NewUserDao(pool)

	// model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	fmt.Println("服務器[新的結構]在8889端口監聽...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客戶端來連接服務器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		go Process(conn)
	}
}
