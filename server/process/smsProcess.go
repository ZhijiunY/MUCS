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

func (sp *SmsProcess) SendGroupMes(mes *message.Message) {

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

		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn, //
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("轉發消息失敗 err =", err)
	}
}
