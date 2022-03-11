package repository

import (
	"context"
	model "demo-go-kit/account/database/model"
	"errors"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
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

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}
	_, err := repo.db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {

	filter := bson.D{
		{"_id", id},
	}
	x, err := repo.filterTasks(filter)
	if err != nil {
		return "", err
	}
	return x[0].Email, nil
}

func (repo *repo) filterTasks(filter interface{}) ([]*model.User, error) {
	var tasks []*model.User
	ctx := context.Background()
	collection := repo.db.Collection("user")
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}
	for cur.Next(ctx) {
		var t model.User
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &t)
	}
	if err := cur.Err(); err != nil {
		return tasks, err
	}
	cur.Close(ctx)
	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}
	return tasks, nil
}
