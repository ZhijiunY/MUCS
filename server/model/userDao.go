package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ZhijiunY/MUCS/common/message"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {

		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (ud *UserDao) Register(user *message.User) (err error) {

	conn := ud.pool.Get()
	if conn == nil {
		return errors.New("failed to get connection from the pool")
	}
	defer conn.Close()

	_, err = ud.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(data))

	if err != nil {
		fmt.Println("保存註冊用戶錯誤 err=", err)
		return
	}
	return
}
