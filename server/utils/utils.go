package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/ZhijiunY/MUCS/common/message"
)

// 這裡將這些方法關聯到結構體中
type Transfer struct {
	// 分析他應該有哪些字段
	Conn net.Conn
	Buf  [8096]byte // 這時傳輸時，使用緩衝
	// 8096
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {

	// buf := make([]byte, 8096)
	fmt.Println("讀取客戶端發送的數據...")
	// conn.Read 在 conn 沒有被關閉的形況下，才會阻塞
	// 如果客戶關閉的 conn 則，就不會組塞
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		// fmt.Println("conn.Read err=", err)
		// 自定義一個錯誤
		// err = errors.New("read pkg header error")
		return
	}

	// 根據 buf[:4] 轉成一個 uint32 類型
	// var pkgLen uint32
	// pkgLen = binary.BigEndian.Uint32(t.Buf[0:4])
	pkgLen := binary.BigEndian.Uint32(t.Buf[0:4])

	// 根據 pkgLen 讀取消息內容
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// fmt.Println("conn.Read fail err=", err)
		// 自定義一個錯誤
		// err = errors.New("read pkg body error")
		return
	}

	// 把 pkgLen 反序列化成 -> message.Message
	// 技術就是一層窗戶紙 &mes!!
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err.Error())
		return
	}
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	if t.Conn == nil {
		return errors.New("nil connection")
	}

	// 先發送一個長度給對方
	// var pkgLen uint32
	// pkgLen = uint32(len(data))
	pkgLen := uint32(len(data))

	// var buf [4]byte
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	// 發送長度
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) fail", err.Error())
		return
	}

	// 發送 data 本身
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err.Error())
		return
	}
	return
}
