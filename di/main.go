package main

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/dig"
	"fmt"
)

type User struct {
	RedisConn redis.Conn
}

func (this *User) GetUsername(id int) (string, error) {

	return redis.String(this.RedisConn.Do("GET", id))
}

func (this *User) SetUsername(id int, username string) bool {
	ok, _ := redis.Bool(this.RedisConn.Do("SET", id, username))
	return ok
}

func NewUser(redisConn redis.Conn) User {
	return User{redisConn}
}

func NewRedis() (redis.Conn, error) {
	return redis.Dial("tcp", "127.0.0.1:6379")
}

func makeDi() *dig.Container {
	di := dig.New()
	di.Provide(NewRedis)
	di.Provide(NewUser)
	return di
}

func main() {
	di := makeDi()

	var username string

	di.Invoke(func(user User) {
		user.SetUsername(1, "jack")
		username, _ = user.GetUsername(1)
	})

	fmt.Print(username)
}
