package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ZhijiunY/MUCS/common/message"
	"github.com/gomodule/redigo/redis"
)

// 我們在服務器啟動後，就初始化一個 userDao 實例
// 把它做成全局的變量，在需要和 redis 操作時，就直接使用即可
var (
	MyUserDao *UserDao
)

// 定義一個 UserDao 結構體
// 完成對 User 結構體的各種操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工廠模式，創建一個 UserDao 實例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 思考一下在 UserDao 應該提供哪些方法給我們
// 1. 根據用戶 id 返回一個 User 實例 +err
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	// 通過給定 id 去 redis 查詢這個用戶
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 錯誤
		if err == redis.ErrNil { // 表示在 users 哈希中，沒有找到對應id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}

	// 這裡我們需要把 res 反序列化成 User 實例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

// 完成登錄校驗 Login
// 1. Login 完成對用戶的驗證
// 2. 如果用戶的 id 和 pwd 都正確，則返回一個 User 實例。
// 3. 如果用戶的 id 或 pwd 有誤，則返回對應的錯誤訊息。
func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	// 先從 UserDao 的連接池中取出一根連接
	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}
	// 這時證明這個用戶是獲取到.
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (ud *UserDao) Register(user *message.User) (err error) {

	// 先從 UserDao 的連接池中取出一根連接
	conn := ud.pool.Get()
	if conn == nil {
		// 處理無法獲取連接的情況
		return errors.New("failed to get connection from the pool")
	}
	defer conn.Close()

	_, err = ud.getUserById(conn, user.UserId)
	if err == nil { // 如果讀取到一個用戶，說明該用戶存在了
		err = ERROR_USER_EXISTS // 用戶存在了
		return
	}

	// 這時，說明 id 在 redis 還沒有，則可以完成註冊
	data, err := json.Marshal(user) // 序列化
	if err != nil {
		return
	}

	// 入庫
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	// _, err = conn.Do("HSet", "users",
	// 	fmt.Sprintf("%d", user.UserId), string(data))

	if err != nil {
		fmt.Println("保存註冊用戶錯誤 err=", err)
		return
	}
	return
}
