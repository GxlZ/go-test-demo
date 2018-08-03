package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	redisConn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	user := NewUser(redisConn)
	user.SetUsername(1)
	username, _ := user.GetUsername(1)
	fmt.Println("username:", username)
}

func NewUser(redisConn redis.Conn) *user {
	return &user{redisConn}
}

type user struct {
	redisConn redis.Conn
}

func (this *user) GetUsername(id int) (string, error) {
	return redis.String(this.redisConn.Do("GET", id))
}

func (this *user) SetUsername(id int) bool {
	ok, _ := redis.Bool(this.redisConn.Do("SET", id, "user_one"))
	return ok
}
