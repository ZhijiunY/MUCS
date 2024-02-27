package message

// 定義一個用戶的結構體
type User struct {
	// 確定字段信息
	// 為了序列化與反序列化成功，我們必須保證
	// 用戶信息的 json 字符串的 key 和 結構體的字段對應的 tag 名字一樣！
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` // 用戶狀態...
	Gender     string `json:"gender"`
}
