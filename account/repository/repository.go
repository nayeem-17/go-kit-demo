package repository

import (
	"context"
	model "demo-go-kit/account/model"
	"errors"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *mongo.Database
	logger log.Logger
}

func NewRepo(db *mongo.Database, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user model.User) error {
	_ = `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	// _, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	// err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	// if err != nil {
	// 	return "", RepoErr
	// }

	return email, nil
}
