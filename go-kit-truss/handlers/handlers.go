package handlers

import (
	"context"

	pb "go-test-demo/go-kit-truss/pb"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UserServer {
	return userService{}
}

type userService struct{}

// GetUsernameV1 implements Service.
func (s userService) GetUsernameV1(ctx context.Context, in *pb.Req) (*pb.Resp, error) {
	var resp pb.Resp
	resp = pb.Resp{
		// Code:
		// Msg:
		// Data:
	}
	return &resp, nil
}
