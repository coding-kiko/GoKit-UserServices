package user

import (
	"context"

	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
	erro "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/errors"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/user/proto"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	getUser    gt.Handler
	createUser gt.Handler
	deleteUser gt.Handler
	updateUser gt.Handler
	proto.UnimplementedUserServicesServer
}

func NewGRPCServer(endpoints Endpoints) proto.UserServicesServer {
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
		deleteUser: gt.NewServer(
			endpoints.DeleteUser,
			decodeDeleteUserReq,
			encodeDeleteUserResp,
		),
		updateUser: gt.NewServer(
			endpoints.UpdateUser,
			decodeUpdateUserReq,
			encodeUpdateUserResp,
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

// decode create user request from outside to endpoints
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

// Encode create response from endpoints to the outside
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

func (s *gRPCServer) DeleteUser(ctx context.Context, req *proto.DeleteUserReq) (*proto.DeleteUserResp, error) {
	_, resp, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		status := erro.ErrToGRPCcode(err)
		resp = &proto.DeleteUserResp{Error: status}
		return resp.(*proto.DeleteUserResp), nil
	}
	return resp.(*proto.DeleteUserResp), nil
}

// decode Delete response coming from 'outside' to the endpoints
func decodeDeleteUserReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.DeleteUserReq)
	return ent.DeleteUserReq{Id: req.Id}, nil
}

// Encode delete response from endpoints to the outside
func encodeDeleteUserResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(ent.DeleteUserResp)
	if !ok { // in case of error
		return &proto.DeleteUserResp{
			Error: &proto.Status{},
		}, nil
	}
	return &proto.DeleteUserResp{
		Deleted: resp.Deleted,
		Error:   &proto.Status{Code: 0, Message: "ok"},
	}, nil
}

func (s *gRPCServer) UpdateUser(ctx context.Context, req *proto.UpdateUserReq) (*proto.UpdateUserResp, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		status := erro.ErrToGRPCcode(err)
		resp = &proto.UpdateUserResp{Error: status}
		return resp.(*proto.UpdateUserResp), nil
	}
	return resp.(*proto.UpdateUserResp), nil
}

// decode update user request from outside to endpoints
func decodeUpdateUserReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.UpdateUserReq)
	return ent.UpdateUserReq{
		Name:        req.Name,
		Age:         req.Age,
		Job:         req.Job,
		Nationality: req.Nationality,
		Pwd:         req.Pwd,
		Email:       req.Email,
	}, nil
}

// Encode update response from endpoints to the outside
func encodeUpdateUserResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(ent.UpdateUserResp)
	if !ok { // in case of error
		return &proto.UpdateUserResp{
			Error: &proto.Status{},
		}, nil
	}
	return &proto.UpdateUserResp{
		Updated: resp.Updated,
		Error:   &proto.Status{Code: 0, Message: "ok"},
	}, nil
}
