package process2

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
	"github.com/ZhijiunY/MUCS/server/utils"
)

type SmsProcess struct {
}

// 寫方法轉發消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) {

	// 遍歷服務器端的onlineUsers map[int]*UserProcess,
	// 將消息轉發出去

	// 取出 mes 的內容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		// 這裏，還需要過濾到自己，即不要再發給自己
		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	// 創建一個 Transfer 實例，發送 data
	tf := &utils.Transfer{
		Conn: conn, //
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("轉發消息失敗 err =", err)
	}
}
