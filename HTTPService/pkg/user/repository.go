package user

import (
	"context"

	erro "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/errors"

	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/user/proto"
	ent "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/entities"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

type gRPCstub struct {
	conn   *grpc.ClientConn
	logger log.Logger
}

func NewGRPClient(log log.Logger, con *grpc.ClientConn) *gRPCstub {
	return &gRPCstub{
		conn:   con,
		logger: log,
	}
}

func (g *gRPCstub) Get(ctx context.Context, req ent.GetUserReq) (ent.GetUserResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.GetUserReq{
		Id: req.Id,
	}
	gRPCresp, err := client.GetUser(ctx, request)
	if err != nil {
		return ent.GetUserResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.GetUserResp{
			Name:        gRPCresp.Name,
			Nationality: gRPCresp.Nationality,
			Job:         gRPCresp.Job,
			Age:         gRPCresp.Age,
			Id:          gRPCresp.Id,
			Created:     gRPCresp.Created,
			Email:       gRPCresp.Email,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.GetUserResp{}, err
	}
}

func (g *gRPCstub) Create(ctx context.Context, req ent.CreateUserReq) (ent.CreateUserResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.CreateUserReq{
		Name:        req.Name,
		Age:         req.Age,
		Nationality: req.Nationality,
		Job:         req.Job,
		Pwd:         req.Pwd,
		Email:       req.Email,
	}
	gRPCresp, err := client.CreateUser(ctx, request)
	if err != nil {
		return ent.CreateUserResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.CreateUserResp{
			Id:      gRPCresp.Id,
			Created: gRPCresp.Created,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.CreateUserResp{}, err
	}
}

func (g *gRPCstub) Delete(ctx context.Context, req ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.DeleteUserReq{
		Id: req.Id,
	}
	gRPCresp, err := client.DeleteUser(ctx, request)
	if err != nil {
		return ent.DeleteUserResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.DeleteUserResp{
			Deleted: gRPCresp.Deleted,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.DeleteUserResp{}, err
	}
}
