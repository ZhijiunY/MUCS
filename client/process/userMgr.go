package process

import (
	"fmt"

	"github.com/ZhijiunY/MUCS/client/model"
	"github.com/ZhijiunY/MUCS/common/message"
)

// 客戶端要維護的 map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser // 我們在用戶登錄成功後，完成對 CurUser 初始化

// 在客戶端顯示當前在線的用戶
func outputOnlineUser() {
	// 遍歷 onlineUsers
	fmt.Println("當前在線用戶列表：")
	// for id, _ := range onlineUsers {
	for id := range onlineUsers {
		// 如果不是，顯示自己
		fmt.Println("用戶id:\t", id)
	}
}

// 編寫一個方法，處理返回的 NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	// 適當優化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { // 原來沒有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
