package model

import (
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
)

// current  當前用戶
type CurUser struct {
	Conn net.Conn
	message.User
}
