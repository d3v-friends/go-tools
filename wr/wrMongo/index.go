package wrMongo

import (
	"context"
	"github.com/d3v-friends/go-tools/fn/fnParams"
	"go.mongodb.org/mongo-driver/mongo"
)

func IndexCollection(
	ctx context.Context,
	colNm string,
	idx []mongo.IndexModel,
	isInits ...bool,
) (_ []string, err error) {
	var col = GetColP(ctx, colNm)

	if fnParams.Get(isInits) {
		if _, err = col.Indexes().DropAll(ctx); err != nil {
			return
		}
	}

	return col.Indexes().CreateMany(ctx, idx)
}
