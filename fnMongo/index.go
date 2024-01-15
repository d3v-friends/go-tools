package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-tools/fnParams"
	"go.mongodb.org/mongo-driver/mongo"
)

func IndexCollection(
	ctx context.Context,
	colNm string,
	idx []mongo.IndexModel,
	isInits ...bool,
) (_ []string, err error) {
	var col = GetCtxColP(ctx, colNm)

	if fnParams.Get(isInits) {
		if _, err = col.Indexes().DropAll(ctx); err != nil {
			return
		}
	}

	return col.Indexes().CreateMany(ctx, idx)
}
