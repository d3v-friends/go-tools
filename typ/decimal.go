package typ

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
	"io"
	"strconv"
)

/*
Decimal:
	model: github.com/d3v-friends/go-tools/typ.Decimal
*/

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
