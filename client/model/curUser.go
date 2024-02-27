package model

import (
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
