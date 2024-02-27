package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/ZhijiunY/MUCS/common/message"
	"github.com/ZhijiunY/MUCS/server/utils"
)

func ShowMenu() {
	fmt.Println("----- 恭喜 xxx 登入成功 -----")
	fmt.Println("----- 1. 顯示在線用戶列表 -----")
	fmt.Println("----- 2. 發送消息 -----")
	fmt.Println("----- 3. 信息列表 -----")
	fmt.Println("----- 4. 退出系統 -----")
	fmt.Println("----- 請選擇(1-4): -----")
	var key int
	var content string

	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("想說的話：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你選擇退出系統...")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确...")
	}
}

// 和服務器端保持通訊
func serverProcessMes(conn net.Conn) {
	// 創建一個 transfer 實例，不停的讀取服務器發送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客戶端正在等待讀取服務器發送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err =", err.Error())
			return
		}
		// 如果讀取到消息，又是下一步處理邏輯
		switch mes.Type {
		case message.NotifyUserStatusMesType: // 有人上線了

			// 1. 取出.NotifyUserStatusMeserr
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 2. 把這個用戶的信息，狀態保存到客戶map[int]User中
			updateUserStatus(&notifyUserStatusMes)
			// 處理
		case message.SmsMesType: // 有人群發消息
			outpurGroupMes(&mes)
		default:
			fmt.Println("服務器端返回了未知的消息類型")
		}
		// fmt.Printf("mes=%v\n", mes)
	}
}
