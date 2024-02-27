package main

import (
	"fmt"
	"io"
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
	process2 "github.com/ZhijiunY/MUCS/server/process"
	"github.com/ZhijiunY/MUCS/server/utils"
)

// 先創建一個
type Processor struct {
	Conn net.Conn
}

func (p *Processor) process2() (err error) {

	// 循環讀客戶端發送的訊息
	for {
		// 這裡我們將讀取數據包，直接封裝成一個函數 readPkg(), 返回 Message, Err
		// 創建一個 Transfer 實例完成讀包任務
		tf := &utils.Transfer{
			Conn: p.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出，服務器端也正常退出...")
				return err
				// continue
			} else {
				fmt.Println("readPkg err=", err.Error())
				return err
			}
		}
		fmt.Println("mes=", mes)

		err = p.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}

// 編寫一個 ServerProcessMes 函數
// 功能：根據客戶端發送消息種類不同，決定用哪個函數來處理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {

	// 看看是否能接收到客戶端發送的群發消息
	// fmt.Println("mes=", mes)

	switch mes.Type {
	case message.LoginMesType: // 如果是登錄的消息
		// 處理登入
		// 創建一個 UserProcess 實例
		up := &process2.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 處理註冊
		up := &process2.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes) // type: data
	case message.SmsMesType:
		//創建一個 SmsProcess 實例完成轉發群聊消息
		smsProcess := &process2.SmsProcess{}
		// msg := string(mes.Data)
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息類型不存在，無法處理...")
	}
	return
}
