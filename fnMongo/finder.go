package fnMongo

import (
	"context"
	"github.com/d3v-friends/go-pure/fnReflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T any](
	ctx context.Context,
	colNm string,
	filter Filter,
	sorter ...Sorter,
) (res *T, err error) {
	var f bson.M
	if f, err = filter.GetFilter(); err != nil {
		return
	}

	var opt = &options.FindOneOptions{}
	if len(sorter) != 0 {
		var s any
		if s, err = sorter[0].GetSorter(); err != nil {
			return
		}
		opt.Sort = s
	}

	var col = GetDBP(ctx).Collection(colNm)
	var cur *mongo.SingleResult
	if cur = col.FindOne(ctx, f, opt); cur.Err() != nil {
		err = cur.Err()
		return
	}

	res = new(T)
	if err = cur.Decode(res); err != nil {
		return
	}

	return
}

func FindAll[T any](
	ctx context.Context,
	colNm string,
	filter Filter,
	sorter ...Sorter,
) (res []*T, err error) {
	var f bson.M
	if f, err = filter.GetFilter(); err != nil {
		return
	}

	var opt = &options.FindOptions{}
	if len(sorter) != 0 {
		var s any
		if s, err = sorter[0].GetSorter(); err != nil {
			return
		}
		opt.Sort = s
	}

	var col = GetDBP(ctx).Collection(colNm)
	var cur *mongo.Cursor
	if cur, err = col.Find(ctx, f, opt); err != nil {
		return
	}

	res = make([]*T, 0)
	if err = cur.All(ctx, &res); err != nil {
		return
	}

	return
}

func FindList[T any](
	ctx context.Context,
	colNm string,
	filter Filter,
	pager Pager,
	sorter ...Sorter,
) (res *ResultList[T], err error) {
	var f bson.M
	if f, err = filter.GetFilter(); err != nil {
		return
	}

	var col = GetDBP(ctx).Collection(colNm)
	var total int64
	if total, err = col.CountDocuments(ctx, f); err != nil {
		return
	}

	var opt = &options.FindOptions{}
	if len(sorter) != 0 {
		var s any
		if s, err = sorter[0].GetSorter(); err != nil {
			return
		}
		opt.Sort = s
	}

	opt.Skip = fnReflect.ToPointer(pager.GetPage() * pager.GetSize())
	opt.Limit = fnReflect.ToPointer(pager.GetSize())

	var cur *mongo.Cursor
	if cur, err = col.Find(ctx, f, opt); err != nil {
		return
	}

	var ls = make([]*T, 0)
	if err = cur.All(ctx, &ls); err != nil {
		return
	}

	res = &ResultList[T]{
		Page:  pager.GetPage(),
		Size:  pager.GetSize(),
		Total: total,
		List:  ls,
	}

	return
}
