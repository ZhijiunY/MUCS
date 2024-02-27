package process2

import "fmt"

// 因為 UserMgr 實例在服務器端有且只有一個
// 因為在很多的地方，都會只用到，因此，我們
// 將其定義為全局變量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成對 userMgr 初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成對 onlineUsers 添加
func (Um *UserMgr) AddOnlineUser(up *UserProcess) {

	Um.onlineUsers[up.UserId] = up
}

// 刪除
func (Um *UserMgr) DelOnlineUser(userId int) {

	delete(Um.onlineUsers, userId)
}

// 查詢：返回當前所有在線的用戶 (獲取所有)
func (Um *UserMgr) GetAllOnlineUser() map[int]*UserProcess { // 返回 map[int]*UserProcess

	return Um.onlineUsers
}

// 根據 id 返回對應的值 （索取單個）
func (Um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	// 如何從 map 取出一個值，帶檢驗方式
	up, ok := Um.onlineUsers[userId]
	if !ok { // 說明，你要查找的這個用戶，當前不再線
		err = fmt.Errorf("用戶 %d 不存在", userId)
		return
	}
	// 取反就為真

	return
}
