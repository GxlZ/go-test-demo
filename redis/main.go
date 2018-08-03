package main

import (
	goredis "github.com/go-redis/redis"
)

func main() {

	redis := NewRedis()

	user := NewUser(redis)

	user.GetUsername()

}

type redis interface {
	getClient() *goredis.Client
}

type client struct {
	options *goredis.Options
}

func (this client) getClient() *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func NewRedis() redis {
	return client{}
}

func NewUser(redis redis) *user {
	return &user{redis}
}

type user struct {
	redis redis
}

func (this *user) GetUsername() string {
	c := this.redis.getClient()
	return c.Get("aaa").String()
}

func (this *user) Set() {

}
