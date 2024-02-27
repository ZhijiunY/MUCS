package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 這裡我們定義幾個用戶狀態的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息類型
	Data string `json:"data"` // 消息內容
}

// 定義兩個消息...後面需要再增加
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用戶id
	UserPwd  string `json:"userPwd"`  // 用戶密碼
	UserName string `json:"userName"` // 用戶名
}

type LoginResMes struct {
	Code    int    `json:"code"`    // 返回狀態碼，500=該用戶未註冊，200=登入成功
	UsersId []int  `json:"usersId"` // 增加字段，保存用戶 id 的切片
	Error   string `json:"error"`   // 返回錯誤消息
}

// 註冊消息
type RegisterMes struct {
	User User `json:"user"` // 類型就是 User 結構體
}

// 註冊響應消息
type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回狀態碼 400 表示該用戶已經佔有 200 表示註冊成功
	Error string `json:"error"` // 返回錯誤訊息
}

// 為了配合服務器推送用戶狀態變化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` // 用戶 id
	Status int `json:"status"` // 用戶的狀態
}

// 增加一個 SmsMes // 發送的消息
type SmsMes struct {
	Content string `json:"content"` // 內容
	User           // 匿名結構體
}

// SmsReMes
