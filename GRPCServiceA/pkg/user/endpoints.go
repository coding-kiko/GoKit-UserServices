package user

import (
	"context"
	"errors"

	ent "github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/entities"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser   endpoint.Endpoint
	GetUser      endpoint.Endpoint
	DeleteUser   endpoint.Endpoint
	UpdateUser   endpoint.Endpoint
	Authenticate endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetUser:      makeGetUserEndpoint(s),
		CreateUser:   makeCreateUserEndpoint(s),
		DeleteUser:   makeDeleteUserEndpoint(s),
		UpdateUser:   makeUpdateUserEndpoint(s),
		Authenticate: makeAuthenticateEndpoint(s),
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ent.GetUserReq)
		if !ok {
			return nil, errors.New("error asserting GetUserReq")
		}
		user, err := s.GetUser(ctx, req)
		if err != nil {
			return user, err
		}
		return user, nil
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ent.CreateUserReq)
		if !ok {
			return nil, errors.New("error asserting CreateUserReq")
		}
		createResp, err := s.CreateUser(ctx, req)
		if err != nil {
			return createResp, err
		}
		return createResp, nil
	}
}

func makeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ent.DeleteUserReq)
		if !ok {
			return nil, errors.New("error asserting DeleteUserReq")
		}
		delResp, err := s.DeleteUser(ctx, req)
		if err != nil {
			return delResp, err
		}
		return delResp, nil
	}
}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ent.UpdateUserReq)
		if !ok {
			return nil, errors.New("error asserting DeleteUserReq")
		}
		resp, err := s.UpdateUser(ctx, req)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}

func makeAuthenticateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ent.AuthenticateReq)
		if !ok {
			return nil, errors.New("error asserting AuthenticateReq")
		}
		resp, err := s.AuthenticateUser(ctx, req)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}
