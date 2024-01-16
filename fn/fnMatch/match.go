package fnMatch

import "fmt"

type (
	FnMatcher[T any] func(v T) bool
)

type OriginType interface {
	bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128 | string
}

func Has[T any](ls []T, matcher FnMatcher[T]) bool {
	for _, v := range ls {
		if matcher(v) {
			return true
		}
	}
	return false
}

func Contain[T OriginType](ls []T, v T) bool {
	for _, item := range ls {
		if v == item {
			return true
		}
	}
	return false
}

func Get[T any](ls []T, matcher FnMatcher[T]) (res T, err error) {
	for _, item := range ls {
		if matcher(item) {
			res = item
			return
		}
	}
	err = fmt.Errorf("not found element")
	return
}

func GetP[T any](ls []T, matcher FnMatcher[T]) (res T) {
	var err error
	if res, err = Get(ls, matcher); err != nil {
		panic(err)
	}
	return
}
