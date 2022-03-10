package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MClient struct {
	Client *mongo.Client
	Ctx    context.Context
}

func (m *MClient) initialize() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb://root:example@localhost:27017/").
		SetServerAPIOptions(serverAPIOptions)
	var err error
	m.Client, err = mongo.Connect(m.Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Client.Ping(m.Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

var mongoClient *MClient

func NewMongoClient() *MClient {
	if mongoClient == nil {
		mongoClient = &MClient{}
		mongoClient.initialize()
		fmt.Println("mongoClient initialized")
	} else {
		fmt.Println("mongoClient already initialized")
	}
	return mongoClient
}
