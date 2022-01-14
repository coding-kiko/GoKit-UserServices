package user

import (
	"context"

	erro "github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/errors"

	"github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/user/proto"
	ent "github.com/coding-kiko/GoKit-Project-Bootcamp/HTTPService/pkg/entities"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

type gRPCstub struct {
	conn   *grpc.ClientConn
	logger log.Logger
}

type Repository interface {
	Get(ctx context.Context, req ent.GetUserReq) (ent.GetUserResp, error)
	Create(ctx context.Context, req ent.CreateUserReq) (ent.CreateUserResp, error)
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
