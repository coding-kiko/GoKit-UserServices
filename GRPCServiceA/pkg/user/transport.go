package user

import (
	"context"

	ent "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/entities"
	erro "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/errors"
	"github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/user/proto"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	getUser    gt.Handler
	createUser gt.Handler
	proto.UnimplementedUserServicesServer
}

func NewGRPCServer(endpoints Endpoints, logger log.Logger) proto.UserServicesServer {
	return &gRPCServer{
		getUser: gt.NewServer(
			endpoints.GetUser,
			decodeGetUserReq,
			encodeGetUserResp,
		),
		createUser: gt.NewServer(
			endpoints.CreateUser,
			decodeCreateUserReq,
			encodeCreateUserResp,
		),
	}
}

// implement the UserServicesServer interface
func (s *gRPCServer) GetUser(ctx context.Context, req *proto.GetUserReq) (*proto.GetUserResp, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		status := erro.ErrToGRPCcode(err)
		resp = &proto.GetUserResp{Error: status}
		return resp.(*proto.GetUserResp), nil
	}
	return resp.(*proto.GetUserResp), nil
}

// decode message coming from 'outside' to the endpoints
func decodeGetUserReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.GetUserReq)
	return ent.GetUserReq{Id: req.Id}, nil
}

// encode message coming from endpoints to the 'outside'
func encodeGetUserResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(ent.GetUserResp)
	if !ok { // in case of error
		return &proto.GetUserResp{
			Error: &proto.Status{},
		}, nil
	}
	return &proto.GetUserResp{
		Id:          resp.Id,
		Name:        resp.Name,
		Job:         resp.Job,
		Nationality: resp.Nationality,
		Created:     resp.Created,
		Age:         resp.Age,
		Email:       resp.Email,
		Error:       &proto.Status{Code: 0, Message: "ok"},
	}, nil
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *proto.CreateUserReq) (*proto.CreateUserResp, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		status := erro.ErrToGRPCcode(err)
		resp = &proto.CreateUserResp{Error: status}
		return resp.(*proto.CreateUserResp), nil
	}
	return resp.(*proto.CreateUserResp), nil
}
func decodeCreateUserReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.CreateUserReq)
	return ent.CreateUserReq{
		Name:        req.Name,
		Age:         req.Age,
		Job:         req.Job,
		Nationality: req.Nationality,
		Pwd:         req.Pwd,
		Email:       req.Email,
	}, nil
}

func encodeCreateUserResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(ent.CreateUserResp)
	if !ok { // in case of error
		return &proto.CreateUserResp{
			Error: &proto.Status{},
		}, nil
	}
	return &proto.CreateUserResp{
		Id:      resp.Id,
		Created: resp.Created,
		Error:   &proto.Status{Code: 0, Message: "ok"},
	}, nil
}
