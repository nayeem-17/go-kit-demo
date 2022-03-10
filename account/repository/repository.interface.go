package repository

import (
	"context"
	model "demo-go-kit/account/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, id string) (string, error)
}
