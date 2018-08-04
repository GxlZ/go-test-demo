package handlers

import (
	"github.com/stretchr/testify/mock"
	"github.com/gomodule/redigo/redis"
	"testing"
	"go-test-demo/go-kit-truss/pb"
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

func TestUserService_GetUsernameV1(t *testing.T) {
	userId := int64(1)
	wantUsername := "jack"

	redisConn := MockRedisConn{}
	redisConn.On(
		"Do",
		"GET",
		userId,
	).Return(wantUsername, nil)

	Convey(fmt.Sprintf("input: &pb.Req{Id:%d}", userId), t, func() {
		userService := NewService(redisConn)
		in := &pb.Req{Id: userId}
		resp, err := userService.GetUsernameV1(nil, in)

		Convey("err is null", func() {
			So(err, ShouldBeEmpty)
		})

		wantCode := 200
		Convey(fmt.Sprintf("resp code: %d", wantCode), func() {
			So(resp.Code, ShouldEqual, 200)
		})

		wantMsg := "success"
		Convey(fmt.Sprintf("resp msg: %s", wantMsg), func() {
			So(resp.Msg, ShouldEqual, wantMsg)
		})

		wantData := "jack"
		Convey(fmt.Sprintf("resp data: %s", wantData), func() {
			So(resp.Data, ShouldEqual, wantData)
		})
	})
}
