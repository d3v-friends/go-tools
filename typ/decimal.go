package typ

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"reflect"
	"strconv"
)

type DecimalCodec struct{}

func DecimalRegistry(registry *bsoncodec.Registry) *bsoncodec.Registry {
	registry.RegisterTypeEncoder(reflect.TypeOf(decimal.Decimal{}), &DecimalCodec{})
	registry.RegisterTypeDecoder(reflect.TypeOf(decimal.Decimal{}), &DecimalCodec{})
	return registry
}

func (dc *DecimalCodec) EncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) (err error) {
	var dec, ok = val.Interface().(decimal.Decimal)
	if !ok {
		err = errors.New("invalid decimal")
		return
	}

	var primitiveDecimal primitive.Decimal128
	if primitiveDecimal, err = primitive.ParseDecimal128(dec.String()); err != nil {
		return
	}

	return vw.WriteDecimal128(primitiveDecimal)
}

func (dc *DecimalCodec) DecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) (err error) {
	var primitiveDecimal primitive.Decimal128
	if primitiveDecimal, err = vr.ReadDecimal128(); err != nil {
		return errors.New("invalid decimal")
	}

	var dec decimal.Decimal
	if dec, err = decimal.NewFromString(primitiveDecimal.String()); err != nil {
		return errors.New("invalid decimal")
	}

	val.Set(reflect.ValueOf(dec))

	return
}

/*------------------------------------------------------------------------------------------------*/
// gqlgen

func MarshalDecimal(b decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b.String())))
	})
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	switch t := v.(type) {
	case string:
		return decimal.NewFromString(t)
	case json.Number:
		return decimal.NewFromString(t.String())
	default:
		var err = fmt.Errorf("invalid Decimal")
		return decimal.Zero, err
	}
}
