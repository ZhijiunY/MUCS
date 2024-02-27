package process

import (
	"encoding/json"

	"github.com/ZhijiunY/MUCS/common/message"

	"fmt"
)

func outpurGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	info := fmt.Sprintf("用戶id:\t%d 對大家說: \t%s",
		smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
