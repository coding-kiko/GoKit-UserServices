package user

import (
	"context"

	ent "github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/entities"
	erro "github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/errors"
	"github.com/coding-kiko/GoKit-UserServices/UserServices/pkg/utils"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
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
	AuthenticateUser(ctx context.Context, r ent.AuthenticateReq) (ent.AuthenticateResp, error)
}

type Repository interface {
	GetUser(ctx context.Context, r ent.GetUserReq) (ent.GetUserResp, error)
	CreateUser(ctx context.Context, user ent.User) (ent.CreateUserResp, error)
	DeleteUser(ctx context.Context, r ent.DeleteUserReq) (ent.DeleteUserResp, error)
	UpdateUser(ctx context.Context, r ent.UpdateUserReq) (ent.UpdateUserResp, error)
	AuthenticateUser(ctx context.Context, r ent.AuthenticateReq) (string, error)
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
	logger := log.With(s.logger, "Service method", "CreateUser")

	// Look for any empty fields
	if utils.CheckEmptyField(r) {
		level.Error(logger).Log("error", "invalid number of arguments")
		return ent.CreateUserResp{}, erro.NewErrInvalidArguments()
	}

	user := ent.User{
		Id:      utils.NewId(),
		Name:    r.Name,
		Age:     r.Age,
		Job:     r.Job,
		Email:   r.Email,
		Country: r.Country,
		Created: utils.TimeNow(),
		PwdHsh:  utils.HashPwd(r.Pwd),
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
	logger := log.With(s.logger, "Service method", "UpdateUser")

	// Look for any empty fields
	if utils.CheckEmptyField(r) {
		level.Error(logger).Log("error", "invalid number of arguments")
		return ent.UpdateUserResp{}, erro.NewErrInvalidArguments()
	}

	r.Pwd = utils.HashPwd(r.Pwd)
	resp, err := s.repo.UpdateUser(ctx, r)
	if err != nil {
		return ent.UpdateUserResp{}, err
	}
	return resp, nil
}

func (s service) AuthenticateUser(ctx context.Context, r ent.AuthenticateReq) (ent.AuthenticateResp, error) {
	logger := log.With(s.logger, "Service method", "AuthenticateUser")

	// Check for any empty fields
	if utils.CheckEmptyField(r) {
		level.Error(logger).Log("error", "invalid number of arguments")
		return ent.AuthenticateResp{}, erro.NewErrInvalidArguments()
	}

	// retrieve pwd from database
	pwdhsh, err := s.repo.AuthenticateUser(ctx, r)
	if err != nil {
		return ent.AuthenticateResp{}, err
	}

	// verify password
	if pwdhsh != utils.HashPwd(r.Pwd) {
		level.Error(logger).Log("error", "incorrect password")
		return ent.AuthenticateResp{}, erro.NewErrInvalidCredentials()
	}

	signedToken := utils.NewToken(r.Email)
	resp := ent.AuthenticateResp{
		Token: signedToken,
	}

	return resp, nil
}
