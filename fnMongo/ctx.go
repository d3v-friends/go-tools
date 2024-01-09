package fnMongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

const ctxClientKey = "CTX_MONGO_CLIENT"
const ctxDbKey = "CTX_MONGO_DB"

func GetClient(ctx context.Context, ctxKeys ...string) (client *mongo.Client, err error) {
	var key = ctxClientKey
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

func GetClientP(ctx context.Context, ctxKeys ...string) (client *mongo.Client) {
	var err error
	if client, err = GetClient(ctx, ctxKeys...); err != nil {
		panic(err)
	}
	return
}

func SetClient(ctx context.Context, client *mongo.Client, ctxKeys ...string) context.Context {
	var key = ctxClientKey
	if 0 < len(ctxKeys) {
		key = ctxKeys[0]
	}
	return context.WithValue(ctx, key, client)
}

func GetDB(ctx context.Context, ctxKeys ...string) (db *mongo.Database, err error) {
	var key = ctxDbKey
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

func GetDBP(ctx context.Context, ctxKeys ...string) (db *mongo.Database) {
	var err error
	if db, err = GetDB(ctx, ctxKeys...); err != nil {
		panic(err)
	}
	return
}

func SetDB(ctx context.Context, db *mongo.Database, ctxKeys ...string) context.Context {
	var key = ctxDbKey
	if 0 < len(ctxKeys) {
		key = ctxKeys[0]
	}
	return context.WithValue(ctx, key, db)
}

func GetCol(ctx context.Context, colNm string) (col *mongo.Collection, err error) {
	var db *mongo.Database
	if db, err = GetDB(ctx); err != nil {
		return
	}
	col = db.Collection(colNm)
	return
}

func GetColP(ctx context.Context, colNm string) (col *mongo.Collection) {
	var err error
	if col, err = GetCol(ctx, colNm); err != nil {
		panic(err)
	}
	return
}
