package user

import (
	"context"

	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/utils"
	"github.com/go-kit/log"
)

type service struct {
	repo   Repository
	logger log.Logger
}

// Service interface describes the user services
type Service interface {
	GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error)
	CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error)
	DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error)
	UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error)
}

type Repository interface {
	GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error)
	CreateUser(ctx context.Context, user ent.User) (ent.CreateUserResp, error)
	DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error)
	UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error)
}

func NewService(logger log.Logger, repository Repository) Service {
	return &service{
		logger: logger,
		repo:   repository,
	}
}

func (s service) GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error) {
	user, err := s.repo.GetUser(ctx, r)
	if err != nil {
		return ent.GetUserResp{}, err
	}
	return user, nil
}

func (s service) CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error) {
	user := ent.User{
		Id:          utils.NewId(),
		Name:        r.Name,
		Age:         r.Age,
		Job:         r.Job,
		Email:       r.Email,
		Nationality: r.Nationality,
		Created:     utils.TimeNow(),
		PwdHsh:      utils.HashPwd(r.Pwd),
	}
	resp, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return ent.CreateUserResp{}, err
	}
	return resp, nil
}

func (s service) DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	resp, err := s.repo.DeleteUser(ctx, r)
	if err != nil {
		return ent.DeleteUserResp{}, err
	}
	return resp, nil
}

func (s service) UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	r.Pwd = utils.HashPwd(r.Pwd)
	resp, err := s.repo.UpdateUser(ctx, r)
	if err != nil {
		return ent.UpdateUserResp{}, err
	}
	return resp, nil
}
