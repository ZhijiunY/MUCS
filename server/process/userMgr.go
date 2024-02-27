package process2

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (Um *UserMgr) AddOnlineUser(up *UserProcess) {

	Um.onlineUsers[up.UserId] = up
}

func (Um *UserMgr) DelOnlineUser(userId int) {

	delete(Um.onlineUsers, userId)
}

func (Um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {

	return Um.onlineUsers
}

func (Um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	up, ok := Um.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用戶 %d 不存在", userId)
		return
	}

	return
}
