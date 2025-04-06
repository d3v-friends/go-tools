package fnPointer

import "reflect"

func Make[T any](v T) *T {
	return &v
}

func IsNil(v any) bool {
	var vo = reflect.ValueOf(v)
	if vo.Kind() != reflect.Pointer {
		return false
	}

	return vo.IsNil()
}

func Default[T any](value *T, defs T) *T {
	if IsNil(value) {
		return &defs
	}
	return value
}
