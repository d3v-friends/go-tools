package mgTree

import (
	"context"
	"github.com/d3v-friends/go-tools/wr/wrMongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// T 는 struct 로 최대한 필드는 자식 단계까지만 사용하는것을 추천

type Model[T any] struct {
	Id        primitive.ObjectID `bson:"_id"`
	ParentId  primitive.ObjectID `bson:"parent_id"`
	Data      T                  `bson:"inline"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

var Migrate = &wrMongo.MigrateModel{
	ColNm: "trees",
	Migrate: []wrMongo.Migrate{
		func(ctx context.Context, col *mongo.Collection) (memo string, err error) {
			memo = "init indexing"
			_, err = col.Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{
							Key:   "parent_id",
							Value: -1,
						},
					},
				},
				{
					Keys: bson.D{
						{
							Key:   "created_at",
							Value: -1,
						},
					},
				},
				{
					Keys: bson.D{
						{
							Key:   "updated_at",
							Value: -1,
						},
					},
				},
			})
			return
		},
	},
}
