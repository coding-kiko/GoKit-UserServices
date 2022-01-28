package user

import (
	"context"

	ent "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/entities"
	"github.com/go-kit/log"
)

type service struct {
	repo   Repository
	logger log.Logger
}
type Service interface {
	GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error)
	CreateUser(ctx context.Context, r ent.CreateUserReq) (ent.CreateUserResp, error)
	DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error)
	UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error)
	AuthenticateUser(ctx context.Context, r ent.AuthenticateReq) (ent.AuthenticateResp, error)
}

type Repository interface {
	Get(ctx context.Context, req ent.GetUserReq) (ent.GetUserResp, error)
	Create(ctx context.Context, req ent.CreateUserReq) (ent.CreateUserResp, error)
	Delete(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error)
	Update(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error)
	Authenticate(ctx context.Context, r ent.AuthenticateReq) (ent.AuthenticateResp, error)
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

// interaction with respository -> grpc stub (client)
func (s service) DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	resp, err := s.repo.Delete(ctx, r)
	if err != nil {
		return ent.DeleteUserResp{}, err
	}
	return resp, nil
}

// interaction with respository -> grpc stub (client)
func (s service) UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	resp, err := s.repo.Update(ctx, r)
	if err != nil {
		return ent.UpdateUserResp{}, err
	}
	return resp, nil
}

// interaction with respository -> grpc stub (client)
func (s service) AuthenticateUser(ctx context.Context, r ent.AuthenticateReq) (ent.AuthenticateResp, error) {
	resp, err := s.repo.Authenticate(ctx, r)
	if err != nil {
		return ent.AuthenticateResp{}, err
	}
	return resp, nil
}
