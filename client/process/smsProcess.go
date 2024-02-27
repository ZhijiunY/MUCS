package process

import (
	"encoding/json"
	"fmt"

	"github.com/ZhijiunY/MUCS/common/message"
	"github.com/ZhijiunY/MUCS/server/utils"
)

type SmsProcess struct {
}

// 發送群聊的消息
func (sp *SmsProcess) SendGroupMes(content string) (err error) {

	// 1. 創建一個 Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. 創建一個 SmsMes 實例
	var smsMes message.SmsMes
	smsMes.Content = content               // 內容
	smsMes.UserId = CurUser.UserId         //
	smsMes.UserStatus = CurUser.UserStatus //

	// 3. 序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	mes.Data = string(data)

	// 4. 對 mes 再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	// 5. 將 mes 發送給服務器..
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	// 6. 發送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}

	return
}
