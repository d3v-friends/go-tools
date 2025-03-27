package fnLogger

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	ErrCannotMarshalMessage = "cannot_marshal_message"
)

func stringify(message any) string {
	var vo = reflect.ValueOf(message)
	if vo.Kind() == reflect.Pointer {
		if !vo.CanInterface() {
			return ""
		}
	}

	switch t := message.(type) {
	case string:
		return t
	case *string:
		return *t
	case error:
		return t.Error()
	case fmt.Stringer:
		return t.String()
	default:
		var body, err = json.Marshal(message)
		if err != nil {
			return ErrCannotMarshalMessage
		}
		return string(body)
	}
}
