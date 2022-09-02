package db

import (

	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"

	"time"
)

type IMongoDb interface {
	GetDb(dbName string) *mongo.Database
	GetContext() context.Context
}

type MongoDb struct {
	client  *mongo.Client
	context context.Context
}

func NewMongoDb() *MongoDb {

		return createLocal()

}
func (mongoDb *MongoDb) GetContext() context.Context {
	return mongoDb.context
}

func createLocal() *MongoDb {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return &MongoDb{
		client:  client,
		context: ctx,
	}
}


func (mongoDb *MongoDb) GetDb(dbName string) *mongo.Database {
	return mongoDb.client.Database(dbName)
}
