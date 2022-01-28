package user

import (
	"context"

	ent "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/entities"
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
		req := request.(ent.GetUserReq)
		user, err := s.GetUser(ctx, req)
		if err != nil {
			return user, err
		}
		return user, nil
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ent.CreateUserReq)
		createResp, err := s.CreateUser(ctx, req)
		if err != nil {
			return createResp, err
		}
		return createResp, nil
	}
}

func makeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ent.DeleteUserReq)
		delResp, err := s.DeleteUser(ctx, req)
		if err != nil {
			return delResp, err
		}
		return delResp, nil
	}
}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ent.UpdateUserReq)
		resp, err := s.UpdateUser(ctx, req)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}

func makeAuthenticateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ent.AuthenticateReq)
		resp, err := s.AuthenticateUser(ctx, req)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}
