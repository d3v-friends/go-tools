package fnSlice

import "fmt"

type (
	FnHas[T any] func(v T) bool
)

func ContainFn[T any](ls []T, fn func(v T) bool) (res bool, err error) {
	if _, err = GetFn(ls, fn); err != nil {
		res = false
		return
	}
	res = true
	return
}

func ContainP[T any](ls []T, fn FnHas[T]) (res bool) {
	var err error
	if res, err = ContainFn(ls, fn); err != nil {
		panic(err)
	}
	return
}

func GetFn[T any](ls []T, fn FnHas[T]) (res T, err error) {
	for _, item := range ls {
		if fn(item) {
			res = item
			return
		}
	}
	err = fmt.Errorf("not found element")
	return
}

func GetP[T any](ls []T, fn FnHas[T]) (res T) {
	var err error
	if res, err = GetFn(ls, fn); err != nil {
		panic(err)
	}
	return
}
