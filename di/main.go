package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/facebookgo/inject"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	redisConn, _ := redis.Dial("tcp", "127.0.0.1:6379")

	var g inject.Graph
	user := user{}
	g.Provide(
		&inject.Object{Value: &user},
		&inject.Object{Value: redisConn},
	)

	user.SetUsername(1, "jack")
	//username, _ := user.GetUsername(1)
	//fmt.Println("username:", username)
}

type user struct {
	RedisConn *redis.Conn `inject:""`
}

func (this *user) GetUsername(id int) (string, error) {

	return redis.String((*this.RedisConn).Do("GET", id))
}

func (this *user) SetUsername(id int, username string) bool {
	spew.Dump(this.RedisConn)
	ok, _ := redis.Bool((*this.RedisConn).Do("SET", id, username))
	return ok
}
