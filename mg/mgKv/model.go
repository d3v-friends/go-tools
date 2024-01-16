package mgKv

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnReflect"
	"github.com/d3v-friends/go-tools/wr/wrMongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Model struct {
	Id        primitive.ObjectID `bson:"_id"`
	Key       string             `bson:"key"`
	Value     string             `bson:"value"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

var KvModel = &wrMongo.MigrateModel{
	ColNm: "kvs",
	Migrate: []wrMongo.Migrate{
		func(ctx context.Context, col *mongo.Collection) (memo string, err error) {
			memo = "init indexing"
			_, err = col.Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{
							Key:   "key",
							Value: 1,
						},
					},
					Options: &options.IndexOptions{
						Unique: fnReflect.ToPointer(true),
					},
				},
			})
			return
		},
	},
}

type KvKey string

func (x KvKey) String() string {
	return string(x)
}

/*------------------------------------------------------------------------------------------------*/

func Get(
	ctx context.Context,
	key fmt.Stringer,
	defaults ...string,
) (res *Model, err error) {
	var now = time.Now()
	var col = wrMongo.GetColP(ctx, KvModel.ColNm)

	var total int64
	if total, err = col.CountDocuments(
		ctx,
		bson.M{
			"key": key.String(),
		},
	); err != nil {
		return
	}

	if total == 0 {
		if len(defaults) == 0 {
			err = fmt.Errorf("not found kv")
			return
		}

		res = &Model{
			Id:        primitive.NewObjectID(),
			Key:       key.String(),
			Value:     defaults[0],
			UpdatedAt: now,
		}

		if _, err = col.InsertOne(ctx, res); err != nil {
			return
		}

		return
	}

	var cur *mongo.SingleResult
	if cur = col.FindOne(
		ctx,
		bson.M{
			"key": key,
		},
	); cur.Err() != nil {
		err = cur.Err()
		return
	}

	res = new(Model)
	if err = cur.Decode(res); err != nil {
		return
	}

	return
}

func Set(
	ctx context.Context,
	key fmt.Stringer,
	value string,
) (err error) {
	_, err = wrMongo.GetColP(ctx, KvModel.ColNm).UpdateOne(
		ctx,
		bson.M{
			"key": key,
		},
		bson.M{
			"$set": bson.M{
				"value":     value,
				"updatedAt": time.Now(),
			},
		},
		&options.UpdateOptions{
			Upsert: fnReflect.ToPointer(true),
		})
	return
}
