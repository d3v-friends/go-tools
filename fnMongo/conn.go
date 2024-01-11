package fnMongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type FnSetRegistry func(registry *bsoncodec.Registry) *bsoncodec.Registry

type ConnectToClient struct {
	Host        string
	Username    string
	Password    string
	SetRegistry []FnSetRegistry
}

type ConnectToDB struct {
	*ConnectToClient
	Database string
}

func createOpt(i *ConnectToClient) (opt *options.ClientOptions) {
	opt = options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s", i.Host)).
		SetReadConcern(readconcern.Majority()).
		SetWriteConcern(writeconcern.Majority()).
		SetAuth(options.Credential{
			Username: i.Username,
			Password: i.Password,
		}).
		SetBSONOptions(&options.BSONOptions{
			UseLocalTimeZone: false,
		})

	opt.Registry = bson.DefaultRegistry

	if len(i.SetRegistry) != 0 {
		for _, registry := range i.SetRegistry {
			opt.Registry = registry(opt.Registry)
		}
	}

	return
}

func ConnectClient(i *ConnectToClient) (client *mongo.Client, err error) {
	var ctx = context.TODO()
	if client, err = mongo.Connect(ctx, createOpt(i)); err != nil {
		return
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}

	return
}

func ConnectDB(i *ConnectToDB) (db *mongo.Database, err error) {
	var client *mongo.Client
	if client, err = ConnectClient(i.ConnectToClient); err != nil {
		return
	}
	db = client.Database(i.Database)
	return
}
