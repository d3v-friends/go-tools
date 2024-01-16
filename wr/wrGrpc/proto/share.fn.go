package proto

import (
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/d3v-friends/go-tools/wr/wrMongo"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*------------------------------------------------------------------------------------------------*/
var timeFormat = time.RFC3339

func NewTime(t time.Time) *Time {
	return &Time{
		Value: t.Format(timeFormat),
	}
}

func (x *Time) Time() (time.Time, error) {
	return time.Parse(timeFormat, x.Value)
}

func (x *Time) TimeP() (res time.Time) {
	var err error
	if res, err = x.Time(); err != nil {
		panic(err)
	}
	return
}

/*------------------------------------------------------------------------------------------------*/

func NewObjectID(p primitive.ObjectID) *ObjectID {
	return &ObjectID{
		Value: p.Hex(),
	}
}

func (x *ObjectID) ObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(x.Value)
}

func (x *ObjectID) ObjectIDP() primitive.ObjectID {
	return fnPanic.Get(x.ObjectID())
}

/*------------------------------------------------------------------------------------------------*/

func NewDecimal(d decimal.Decimal) *Decimal {
	return &Decimal{
		Value: d.String(),
	}
}

func (x *Decimal) Decimal() (decimal.Decimal, error) {
	return decimal.NewFromString(x.Value)
}

func (x *Decimal) DecimalP() decimal.Decimal {
	return fnPanic.Get(x.Decimal())
}

/*------------------------------------------------------------------------------------------------*/

func (x *Period) Query() (res bson.M, err error) {
	res = make(bson.M)

	if x.Equal != nil {
		if res["$equal"], err = x.Equal.Time(); err != nil {
			return
		}
	}

	if x.NotEqual != nil {
		if res["$ne"], err = x.NotEqual.Time(); err != nil {
			return
		}
	}

	if x.Gt != nil {
		if res["$gt"], err = x.Gt.Time(); err != nil {
			return
		}
	}

	if x.Gte != nil {
		if res["$gte"], err = x.Gte.Time(); err != nil {
			return
		}
	}

	if x.Lt != nil {
		if res["$lt"], err = x.Lt.Time(); err != nil {
			return
		}
	}

	if x.Lte != nil {
		if res["$lte"], err = x.Lte.Time(); err != nil {
			return
		}
	}

	return
}

/*------------------------------------------------------------------------------------------------*/

func (x IdxDir) IdxDir() wrMongo.IdxDir {
	switch x.String() {
	case IdxDir_ASC.String():
		return 1
	case IdxDir_DESC.String():
		return -1
	default:
		var err = fmt.Errorf("invalid sort value")
		panic(err)
	}
}
