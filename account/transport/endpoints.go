package server

import (
	"context"
	service "demo-go-kit/account/controller"
	dto "demo-go-kit/account/converter"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return dto.CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return dto.GetUserResponse{
			Email: email,
		}, err
	}
}
