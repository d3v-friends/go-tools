package typ

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"strconv"
)

/*
ObjectID:
	model: github.com/d3v-friends/go-tools/typ.ObjectID
*/

func MarshalObjectID(b primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(b.Hex())))
	})
}

func UnmarshalObjectID(v interface{}) (primitive.ObjectID, error) {
	switch t := v.(type) {
	case string:
		return primitive.ObjectIDFromHex(t)
	case []byte:
		return primitive.ObjectIDFromHex(string(t))
	default:
		var err = fmt.Errorf("invalid object id: value=%s", t)
		return primitive.NilObjectID, err
	}
}
