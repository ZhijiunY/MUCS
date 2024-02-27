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

// func init() {
// 	// 當服務器啟動時，我們就去初始化我們的 redis 的連接池
// 	initPool("localhost:6379", 16, 0, 300*time.Second)

// 	initUserDao()
// }

func main() {

	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	// 提示信息
	fmt.Println("服務器[新的結構]在8889端口監聽...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	// 一但監聽成功，就等待客戶端來連接服務器
	for {
		fmt.Println("等待客戶端來連接服務器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		// 一但連接成功，則啟動一個協成和客戶保持通訊...
		go Process(conn)
	}

}

// cd path/to/your/go/project
// /Users/yzj90596/go/bin/dlv debug
