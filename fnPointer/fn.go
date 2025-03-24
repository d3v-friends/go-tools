package fnPointer

import "reflect"

func Make[T any](v T) *T {
	return &v
}

func IsNil(v any) bool {
	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		return false
	}

	var vo = reflect.ValueOf(v)
	return vo.IsNil()
}

func Default[T any](value *T, defs T) *T {
	if IsNil(value) {
		return &defs
	}
	return value
}
