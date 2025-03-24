package fnLogger

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func stringify(message any) string {
	var vo = reflect.ValueOf(message)
	if vo.Kind() == reflect.Pointer {
		if !vo.CanInterface() {
			return ""
		}
	}

	var stringer, ok = message.(fmt.Stringer)
	if ok {
		return stringer.String()
	}

	switch t := message.(type) {
	case string:
		return t
	case *string:
		return *t
	case error:
		return t.Error()
	default:
		var body, err = json.Marshal(message)
		if err != nil {
			return "cannot_marshal_message"
		}
		return string(body)
	}
}
