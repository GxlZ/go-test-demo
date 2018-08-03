package main

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

type MockRedisConn struct {
	mock.Mock
	redis.Conn
}

func (this MockRedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	argsList := []interface{}{commandName}
	argsList = append(argsList, args...)
	mockArgs := this.Mock.Called(argsList...)
	return mockArgs.Get(0), mockArgs.Error(1)
}

func TestGetUsername(t *testing.T) {
	userId := 1
	wantUsername := "jack"

	redisConn := MockRedisConn{}
	redisConn.On(
		"Do",
		"GET",
		userId,
	).Return(wantUsername, nil)

	Convey(fmt.Sprintf("user.name get test, id:%d", userId), t, func() {
		user := NewUser(&redisConn)
		username, _ := user.GetUsername(userId)
		So(username, ShouldEqual, wantUsername)
	})
}
