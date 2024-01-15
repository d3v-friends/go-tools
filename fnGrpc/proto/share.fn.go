package proto

import (
	"github.com/d3v-friends/go-pure/fnPanic"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*------------------------------------------------------------------------------------------------*/
var timeFormat = time.RFC3339

func NewTime(t time.Time) Time {
	return Time{
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

func NewObjectID(p primitive.ObjectID) ObjectID {
	return ObjectID{
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

func NewDecimal(d decimal.Decimal) Decimal {
	return Decimal{
		Value: d.String(),
	}
}

func (x *Decimal) Decimal() (decimal.Decimal, error) {
	return decimal.NewFromString(x.Value)
}

func (x *Decimal) DecimalP() decimal.Decimal {
	return fnPanic.Get(x.Decimal())
}
