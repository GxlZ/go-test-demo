package main

import (
	"github.com/go-redis/redis"
)
func main() {

}

func NewUser(redisClient *goredis.Client) *user {
	return &user{}
}

type user struct {
}

func (this *user) Get() {

}

func (this *user) Set() {

}
