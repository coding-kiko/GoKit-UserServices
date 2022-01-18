package user

import (
	"context"

	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
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
}

type Repository interface {
	GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error)
	CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error)
	DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error)
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
	resp, err := s.repo.CreateUser(ctx, r)
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
