package typ

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"strconv"
)

func MarshalObjectID(b primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b.Hex())))
	})
}

func UnmarshalObjectID(v interface{}) (res primitive.ObjectID, err error) {
	switch t := v.(type) {
	case string:
		return primitive.ObjectIDFromHex(t)
	case []byte:
		return primitive.ObjectIDFromHex(string(t))
	default:
		err = fmt.Errorf("invalid ObjectID")
		return
	}
}
