package fnPointer

import "reflect"

func IsNil(i any) bool {
	if i == nil {
		return true
	}

	var v = reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func Default[T any](value *T, defs T) T {
	if IsNil(value) {
		return defs
	}
	return *value
}
