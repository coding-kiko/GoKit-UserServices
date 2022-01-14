package user

import (
	"context"

	ent "github.com/coding-kiko/GoKit-Project-Bootcamp/HTTPService/pkg/entities"
	"github.com/go-kit/kit/log"
)

type service struct {
	repo   Repository
	logger log.Logger
}
type Service interface {
	GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error)
	CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error)
}

func NewService(log log.Logger, repository Repository) Service {
	return &service{
		logger: log,
		repo:   repository,
	}
}

// interaction with respository -> grpc stub (client)
func (s service) GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error) {
	user, err := s.repo.Get(ctx, r)
	if err != nil {
		return ent.GetUserResp{}, err
	}
	return user, nil
}

// interaction with respository -> grpc stub (client)
func (s service) CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error) {
	resp, err := s.repo.Create(ctx, r)
	if err != nil {
		return ent.CreateUserResp{}, err
	}
	return resp, nil
}
