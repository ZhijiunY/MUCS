package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/ZhijiunY/MUCS/common/message"
	"github.com/ZhijiunY/MUCS/server/utils"
)

type UserProcess struct {
}

func (up *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err.Error())
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("註冊發送消息錯誤 err=", err.Error())
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err.Error())
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("註冊成功，請重新登入")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

func (up *UserProcess) Login(userId int, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	mes := message.Message{
		Type: message.LoginMesType,
	}

	loginMes := message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	pkgLen := uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) fail", err)
		return
	}

	fmt.Printf("客戶端，發送訊息的長度=%d 內容=%s", len(data), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write(buf) fail", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		fmt.Println("當前在線用戶列表如下：")
		for _, v := range loginResMes.UsersId {

			if v == userId {
				continue
			}

			fmt.Println("用戶 id: \t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}

			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
