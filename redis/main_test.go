package main

import (
	"testing"
	goredis "github.com/go-redis/redis"
	"github.com/stretchr/testify/mock"
	"log"
)

type MockRedis struct {
	mock.Mock
}

func (this MockRedis) getClient() *goredis.Client {
	args := this.Mock.Called()
	return args.Get(0).(*goredis.Client)
}

func TestUserGet(t *testing.T) {
	redis := MockRedis{}
	redis.On("getClient",
	).Return(goredis.NewClient(&goredis.Options{}))

	user := NewUser(redis)
	username := user.GetUsername()

	log.Println(username)
}
