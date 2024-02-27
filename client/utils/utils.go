package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("讀取客戶端發送的數據...")

	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {

		return
	}

	pkgLen := uint32(binary.BigEndian.Uint32(t.Buf[0:4]))
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {

		return
	}

	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err.Error())
		return
	}
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {

	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) fail", err.Error())
		return
	}

	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err.Error())
		return
	}
	return
}
