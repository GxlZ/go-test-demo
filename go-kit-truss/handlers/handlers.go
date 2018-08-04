package handlers

import (
	"context"

	"github.com/gomodule/redigo/redis"
	"go-test-demo/go-kit-truss/pb"
)

func NewService(redisConn redis.Conn) UserService {
	return UserService{redisConn}
}

type UserService struct {
	RedisConn redis.Conn
}

// GetUsernameV1 implements Service.
func (s UserService) GetUsernameV1(ctx context.Context, in *pb.Req) (*pb.Resp, error) {
	username, _ := redis.String(s.RedisConn.Do("GET", in.Id))

	var resp pb.Resp
	resp = pb.Resp{
		Code: 200,
		Msg:  "success",
		Data: username,
	}
	return &resp, nil
}
