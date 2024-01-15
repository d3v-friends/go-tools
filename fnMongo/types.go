package fnMongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Pager interface {
		GetSize() int64
		GetPage() int64
	}

	Indexer interface {
		GetIndex() []mongo.IndexModel
	}

	Filter interface {
		GetFilter() (bson.M, error)
		GetColNm() string
	}

	Sorter interface {
		GetSorter() (bson.M, error)
	}

	Updater interface {
		GetUpdater() (bson.M, error)
	}

	ResultList[T any] struct {
		Page  int64
		Size  int64
		Total int64
		List  []*T
	}
)

/*------------------------------------------------------------------------------------------------*/

type IdxDir int

func (x IdxDir) Int() int {
	return int(x)
}

func (x IdxDir) Valid() bool {
	for _, dir := range IdxDirAll {
		if x == dir {
			return true
		}
	}
	return false
}

var (
	IdxDirASC  IdxDir = 1
	IdxDirDESC IdxDir = -1
)

var IdxDirAll = []IdxDir{
	IdxDirASC,
	IdxDirDESC,
}

/*------------------------------------------------------------------------------------------------*/
