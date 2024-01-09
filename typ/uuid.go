package typ

import (
	"database/sql/driver"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"io"
	"reflect"
	"strconv"
)

type UUID string

func NewUUID() UUID {
	return UUID(uuid.NewString())
}

func (x *UUID) Scan(src any) (err error) {
	switch v := src.(type) {
	case string:
		*x = UUID(v)
		return
	case []byte:
		*x = UUID(v)
		return
	default:
		err = fmt.Errorf("invalid src type")
		return
	}
}

func (x UUID) Value() (res driver.Value, err error) {
	res = string(x)
	return
}

func (x *UUID) String() string {
	return string(*x)
}

/*------------------------------------------------------------------------------------------------*/
// for gqlgen

func MarshalUUID(b UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b.String())))
	})
}

func UnmarshalUUID(v any) (res UUID, err error) {
	switch t := v.(type) {
	case string:
		res = UUID(t)
		return
	case []byte:
		res = UUID(t)
		return
	default:
		err = fmt.Errorf("invalid UUID scalar")
		return
	}
}

/*------------------------------------------------------------------------------------------------*/
// for mongo-driver

type UUIDCodec struct{}

func UUIDRegistry(registry *bsoncodec.Registry) *bsoncodec.Registry {
	registry.RegisterTypeEncoder(reflect.TypeOf(UUID("")), &UUIDCodec{})
	registry.RegisterTypeDecoder(reflect.TypeOf(UUID("")), &UUIDCodec{})
	return registry
}

func (x *UUIDCodec) EncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) (err error) {
	var v, isOk = val.Interface().(UUID)
	if !isOk {
		err = fmt.Errorf("invalid UUID type")
		return
	}
	return vw.WriteString(v.String())
}

func (x *UUIDCodec) DecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) (err error) {
	var str string
	if str, err = vr.ReadString(); err != nil {
		return
	}
	val.Set(reflect.ValueOf(UUID(str)))
	return
}
