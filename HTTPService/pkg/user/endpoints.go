package user

import (
	"context"

	ent "github.com/coding-kiko/GoKit-Project-Bootcamp/HTTPService/pkg/entities"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetUser:    makeGetUserEndpoint(s),
		CreateUser: makeCreateUserEndpoint(s),
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
