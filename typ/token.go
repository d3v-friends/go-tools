package typ

import (
	"database/sql/driver"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

type Token string

func (x *Token) String() string {
	return string(*x)
}

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