package main

import (
	"fmt"
	"io"
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
	process2 "github.com/ZhijiunY/MUCS/server/process"
	"github.com/ZhijiunY/MUCS/server/utils"
)

type Processor struct {
	Conn net.Conn
}

func (p *Processor) process2() (err error) {

	for {
		tf := &utils.Transfer{
			Conn: p.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出，服務器端也正常退出...")
				return err
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

func (p *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息類型不存在，無法處理...")
	}
	return
}
