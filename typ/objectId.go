package typ

import (
	"database/sql/driver"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"reflect"
	"strconv"
)

type ObjectID primitive.ObjectID

var NilObjectID = ObjectID(primitive.NilObjectID)

func NewObjectID() ObjectID {
	return ObjectID(primitive.NewObjectID())
}

func (x *ObjectID) Scan(src any) (err error) {
	switch v := src.(type) {
	case string:
		if *x, err = ObjectIDFromHexString(v); err != nil {
			return
		}
		return
	case []byte:
		if *x, err = ObjectIDFromHexString(string(v)); err != nil {
			return
		}
		return
	default:
		err = fmt.Errorf("invalid objectId type")
		return
	}
}

func (x ObjectID) Value() (res driver.Value, err error) {
	res = x.String()
	return
}

func (x *ObjectID) String() string {
	return x.ObjectID().Hex()
}

func (x *ObjectID) ObjectID() primitive.ObjectID {
	return primitive.ObjectID(*x)
}

func ObjectIDFromHexString(str string) (objId ObjectID, err error) {
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(str); err != nil {
		return
	}
	objId = ObjectID(id)
	return
}

/*------------------------------------------------------------------------------------------------*/
// mongo codec

type ObjectIDCodec struct {
}

func ObjectIDRegistry(registry *bsoncodec.Registry) *bsoncodec.Registry {
	registry.RegisterTypeEncoder(reflect.TypeOf(ObjectID{}), &ObjectIDCodec{})
	registry.RegisterTypeDecoder(reflect.TypeOf(ObjectID{}), &ObjectIDCodec{})
	return registry
}

func (x *ObjectIDCodec) EncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) (err error) {
	var id, ok = val.Interface().(ObjectID)
	if !ok {
		err = fmt.Errorf("invalid object id")
		return
	}
	return vw.WriteObjectID(id.ObjectID())
}

func (x *ObjectIDCodec) DecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) (err error) {
	var id primitive.ObjectID
	if id, err = vr.ReadObjectID(); err != nil {
		return
	}
	val.Set(reflect.ValueOf(ObjectID(id)))
	return
}

/*------------------------------------------------------------------------------------------------*/
// gqlgen

func MarshalObjectID(b ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b.String())))
	})
}

func UnmarshalObjectID(v interface{}) (ObjectID, error) {
	switch t := v.(type) {
	case string:
		return ObjectIDFromHexString(t)
	case []byte:
		return ObjectIDFromHexString(string(t))
	default:
		var err = fmt.Errorf("invalid ObjectID")
		return ObjectID(primitive.NilObjectID), err
	}
}
