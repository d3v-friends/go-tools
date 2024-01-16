package wrMongo

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

type CodecRegister func(*bsoncodec.Registry) *bsoncodec.Registry

type ConnectClientArgs struct {
	Host           string
	Username       string
	Password       string
	CodecRegisters []CodecRegister
}

func createOpt(i *ConnectClientArgs) (opt *options.ClientOptions) {
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

	if len(i.CodecRegisters) != 0 {
		for _, registry := range i.CodecRegisters {
			opt.Registry = registry(opt.Registry)
		}
	}

	return
}

func ConnectClient(i *ConnectClientArgs) (client *mongo.Client, err error) {
	var ctx = context.TODO()
	if client, err = mongo.Connect(ctx, createOpt(i)); err != nil {
		return
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}

	return
}

/*------------------------------------------------------------------------------------------------*/

type ConnectDBArgs struct {
	*ConnectClientArgs
	Database string
}

func ConnectDB(i *ConnectDBArgs) (db *mongo.Database, err error) {
	var client *mongo.Client
	if client, err = ConnectClient(i.ConnectClientArgs); err != nil {
		return
	}
	db = client.Database(i.Database)
	return
}
