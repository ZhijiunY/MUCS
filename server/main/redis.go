package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// 定義一個全局的 pool
var pool *redis.Pool

func initPool(addr string, idleConn, maxConn int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     idleConn,    // 最大空閒連接數
		MaxActive:   maxConn,     // 表示和數據庫的最大鏈接數，0表示沒有限制
		IdleTimeout: idleTimeout, // 最大空閒時間
		Dial: func() (redis.Conn, error) { // 初始化鏈接的代碼，鏈接哪個 IP 的 redis
			return redis.Dial("tcp", addr)
		},
	}
}
