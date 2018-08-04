package di

import (
	"github.com/gomodule/redigo/redis"
)

func NewRedisConn() (redis.Conn, error) {
	return redis.Dial("tcp", "127.0.0.1:6379")
}
