package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-pure/fnReflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MigrateArgs struct {
	Models []*MigrateModel
}

func Migrate(
	ctx context.Context,
	i *MigrateArgs,
) (err error) {
	var modelList = make([]*MigrateModel, 0)
	modelList = append(modelList, &MigrateModel{
		ColNm:   ColNmMango,
		Migrate: MigrateMango,
	})
	modelList = append(modelList, i.Models...)

	var db = GetDBP(ctx)
	var colMango = db.Collection(modelList[0].ColNm)
	var now = time.Now()
	for _, model := range modelList {
		var count int64
		if count, err = colMango.CountDocuments(ctx, bson.M{
			"colNm": model.ColNm,
		}); err != nil {
			return
		}

		if count == 0 {
			if _, err = colMango.InsertOne(
				ctx,
				&Mango{
					Id:        primitive.NewObjectID(),
					ColNm:     model.ColNm,
					NextIdx:   0,
					History:   make([]*MangoHistory, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
			); err != nil {
				return
			}
		}

		var cur *mongo.SingleResult
		if cur = colMango.FindOne(
			ctx,
			bson.M{
				"colNm": model.ColNm,
			},
		); cur.Err() != nil {
			err = cur.Err()
			return
		}

		var doc = new(Mango)
		if err = cur.Decode(doc); err != nil {
			return
		}

		var colModel = db.Collection(model.ColNm)
		var migrateList = model.Migrate

		for i := doc.NextIdx; i < len(migrateList); i++ {
			var fn = migrateList[i]
			var memo string
			if memo, err = fn(ctx, colModel); err != nil {
				return
			}

			if _, err = colMango.UpdateOne(
				ctx,
				bson.M{
					"colNm": model.ColNm,
				},
				bson.M{
					"$push": bson.M{
						"history": &MangoHistory{
							Memo:       memo,
							MigratedAt: time.Now(),
						},
					},
					"$inc": bson.M{
						"nextIdx": 1,
					},
				}); err != nil {
				return err
			}
		}
	}

	return
}

/*------------------------------------------------------------------------------------------------*/

type (
	Mango struct {
		Id        primitive.ObjectID `bson:"_id"`
		ColNm     string             `bson:"colNm"`
		NextIdx   int                `bson:"nextIdx"`
		History   []*MangoHistory    `bson:"history"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	MangoHistory struct {
		Memo       string    `bson:"memo"`
		MigratedAt time.Time `bson:"migratedAt"`
	}
)

const ColNmMango = "mango"

var MigrateMango = []FnMigrate{
	func(ctx context.Context, col *mongo.Collection) (memo string, err error) {
		memo = "init indexing"
		_, err = col.Indexes().CreateMany(ctx, []mongo.IndexModel{
			{
				Keys: bson.D{
					{
						Key:   "colNm",
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
}
