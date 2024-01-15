package fnMongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

const CtxMongoClientKey = "CTX_MONGO_CLIENT"
const CtxMongoDbKey = "CTX_MONGO_DB"

func GetCtxClient(ctx context.Context, ctxKeys ...string) (client *mongo.Client, err error) {
	var key = CtxMongoClientKey
	if 0 < len(ctxKeys) {
		key = ctxKeys[0]
	}

	var has bool
	if client, has = ctx.Value(key).(*mongo.Client); !has {
		err = fmt.Errorf("not found *mongo.Client")
		return
	}

	return
}

func GetCtxClientP(ctx context.Context, ctxKeys ...string) (client *mongo.Client) {
	var err error
	if client, err = GetCtxClient(ctx, ctxKeys...); err != nil {
		panic(err)
	}
	return
}

func SetCtxClient(ctx context.Context, client *mongo.Client, ctxKeys ...string) context.Context {
	var key = CtxMongoClientKey
	if 0 < len(ctxKeys) {
		key = ctxKeys[0]
	}
	return context.WithValue(ctx, key, client)
}

func GetCtxDB(ctx context.Context, ctxKeys ...string) (db *mongo.Database, err error) {
	var key = CtxMongoDbKey
	if 0 < len(ctxKeys) {
		key = ctxKeys[0]
	}

	var has bool
	if db, has = ctx.Value(key).(*mongo.Database); !has {
		err = fmt.Errorf("not found *mongo.Database")
		return
	}

	return
}

func GetCtxDBP(ctx context.Context, ctxKeys ...string) (db *mongo.Database) {
	var err error
	if db, err = GetCtxDB(ctx, ctxKeys...); err != nil {
		panic(err)
	}
	return
}

func SetCtxDB(ctx context.Context, db *mongo.Database, ctxKeys ...string) context.Context {
	var key = CtxMongoDbKey
	if 0 < len(ctxKeys) {
		key = ctxKeys[0]
	}
	return context.WithValue(ctx, key, db)
}

func GetCtxCol(ctx context.Context, colNm string) (col *mongo.Collection, err error) {
	var db *mongo.Database
	if db, err = GetCtxDB(ctx); err != nil {
		return
	}
	col = db.Collection(colNm)
	return
}

func GetCtxColP(ctx context.Context, colNm string) (col *mongo.Collection) {
	var err error
	if col, err = GetCtxCol(ctx, colNm); err != nil {
		panic(err)
	}
	return
}
