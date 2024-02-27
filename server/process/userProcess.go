package process2

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
	"github.com/ZhijiunY/MUCS/server/model"
	"github.com/ZhijiunY/MUCS/server/utils"
)

type UserProcess struct {
	// 字段
	Conn net.Conn
	// 增加一個字段，表示該 Conn 是哪個用戶
	UserId int
}

// 這裡我們編寫通知所有在線的用戶的方法
// userId 要通知其他的在線用戶，我上線
func (Up *UserProcess) NotifyOthersOnlineUser(userId int) {

	// 遍歷 onlineUsers, 然後一個一個的發送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		// 過濾到自己
		if id == userId {
			continue
		}
		// 開始通知「單獨的一個寫法」
		up.NotifyMeOnline(userId)
	}
}

func (Up *UserProcess) NotifyMeOnline(userId int) {

	// 組裝我們的 NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// 將 notifyUserStatusMes 序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 將序列化後的 notifyUserStatusMes 賦值給 mes.Data
	mes.Data = string(data)

	// 對 mes 再次序列化，準備發送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 發送，創建我們 Transfer 實例，發送
	tf := &utils.Transfer{
		Conn: Up.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err =", err)
		return
	}
}

func (Up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	// 1. 先從 mes 中取出 mes.Data，並直接反序列化成 RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}

	// (1) 先聲明一個 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// 我們需要到 redis 數據庫完成註冊。
	// 1. 使用 model.MyUserDao 到 redis 去驗證
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "註冊發生未知錯誤..."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// (4) 將 data 賦值給 resMes
	resMes.Data = string(data)

	// (5) 對 resMes 逕行序列化，準備發送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// (6) 發送 data ，將其封裝到 writerPkg
	// 因為使用分層模式(mvc)，我們先創建一個Transfer 實例，然後讀取
	tf := &utils.Transfer{
		Conn: Up.Conn,
	}
	err = tf.WritePkg(data)
	return

}

// 編寫一個 ServerProcessLogin 函數，專門處理登入請求
func (Up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代碼...
	// 1. 先從 mes 中取出 mes.Data，並直接反序列化成 LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}

	// (1) 先聲明一個 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// (2) 再聲明一個 LoginResMes，並完成賦值
	var loginResMes message.LoginResMes

	// 我們需要到 redis 數據庫完成驗證。
	// 1. 使用 model.MyUserDao 到 redis 去驗證
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服務器內部錯誤..."
		}
	} else {
		loginResMes.Code = 200
		// 這裡，因為用戶登入成功，我們就把該登入成功的用戶放到 userMgr 中
		// 將登入成功的用戶的 userId 賦給 Up
		Up.UserId = loginMes.UserId
		userMgr.AddOnlineUser(Up)
		// 通知其他的在線用戶，我上線了
		Up.NotifyOthersOnlineUser(loginMes.UserId)
		// 當前在線用戶的 id 放入到 loginResMes.UsersId
		// 遍歷 userMgr.onlineUsers

		// for id, _ := range userMgr.onlineUsers {
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登入成功")
	}

	// // 如果用戶 id = 100,  密碼 = 123456, 認為合法，否則不合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	// 合法
	// 	loginResMes.Code = 200
	// } else {
	// 	// 不合法
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "該用戶不存在，請註冊再使用..."
	// }

	// (3) 將 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// (4) 將 data 賦值給 resMes
	resMes.Data = string(data)

	// (5) 對 resMes 逕行序列化，準備發送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// (6) 發送 data ，將其封裝到 writerPkg
	// 因為使用分層模式(mvc)，我們先創建一個Transfer 實例，然後讀取
	tf := &utils.Transfer{
		Conn: Up.Conn,
	}
	err = tf.WritePkg(data)
	return
}
