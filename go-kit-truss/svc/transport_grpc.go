// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 7dc4d5d85c
// Version Date: 2018年 5月28日 星期一 22时12分59秒 UTC

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "go-test-demo/go-kit-truss/pb"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC UserServer.
func MakeGRPCServer(endpoints Endpoints) pb.UserServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	return &grpcServer{
		// user

		getusernamev1: grpctransport.NewServer(
			endpoints.GetUsernameV1Endpoint,
			DecodeGRPCGetUsernameV1Request,
			EncodeGRPCGetUsernameV1Response,
			serverOptions...,
		),
	}
}

// grpcServer implements the UserServer interface
type grpcServer struct {
	getusernamev1 grpctransport.Handler
}

// Methods for grpcServer to implement UserServer interface

func (s *grpcServer) GetUsernameV1(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	_, rep, err := s.getusernamev1.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Resp), nil
}

// Server Decode

// DecodeGRPCGetUsernameV1Request is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getusernamev1 request to a user-domain getusernamev1 request. Primarily useful in a server.
func DecodeGRPCGetUsernameV1Request(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Req)
	return req, nil
}

// Server Encode

// EncodeGRPCGetUsernameV1Response is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getusernamev1 response to a gRPC getusernamev1 reply. Primarily useful in a server.
func EncodeGRPCGetUsernameV1Response(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Resp)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
