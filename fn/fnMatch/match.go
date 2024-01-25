package fnMatch

import "fmt"

type (
	FnMatcher[T any] func(v T) bool
)

func Has[T any](ls []T, matcher FnMatcher[T]) bool {
	for _, v := range ls {
		if matcher(v) {
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
