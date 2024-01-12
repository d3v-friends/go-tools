package typ

import (
	"database/sql/driver"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"io"
	"reflect"
	"strconv"
)

type Token string

func (x *Token) String() string {
	return string(*x)
}

/*------------------------------------------------------------------------------------------------*/
// gorm

func (x *Token) Scan(src any) (err error) {
	switch v := src.(type) {
	case string:
		*x = Token(v)
		return
	case []byte:
		*x = Token(v)
		return
	default:
		err = fmt.Errorf("invalid src type")
		return
	}
}

func (x Token) Value() (res driver.Value, err error) {
	res = string(x)
	return
}

/*------------------------------------------------------------------------------------------------*/
// gqlgen

func MarshalToken(b string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b)))
	})
}

func UnmarshalToken(v interface{}) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	case []byte:
		return string(t), nil
	default:
		var err = fmt.Errorf("invalid Token scalar")
		return "", err
	}
}

/*------------------------------------------------------------------------------------------------*/
// mongo codec

type TokenCodec struct{}

func TokenRegistry(reg *bsoncodec.Registry) *bsoncodec.Registry {
	reg.RegisterTypeEncoder(reflect.TypeOf(Token("")), &TokenCodec{})
	reg.RegisterTypeDecoder(reflect.TypeOf(Token("")), &TokenCodec{})
	return reg
}

func (x *TokenCodec) EncodeValue(_ bsoncodec.EncodeContext, writer bsonrw.ValueWriter, value reflect.Value) (err error) {
	var v, isOk = value.Interface().(Token)
	if !isOk {
		err = fmt.Errorf("invalid token value")
		return
	}
	return writer.WriteString(v.String())
}

func (x *TokenCodec) DecodeValue(_ bsoncodec.DecodeContext, reader bsonrw.ValueReader, value reflect.Value) (err error) {
	var str string
	if str, err = reader.ReadString(); err != nil {
		return
	}
	value.Set(reflect.ValueOf(Token(str)))
	return
}
