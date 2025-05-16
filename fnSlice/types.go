package fnSlice

import (
	"github.com/d3v-friends/go-tools/fnError"
	"reflect"
)

type List[T any] []T

func (x List[T]) Find(fn func(v T) bool) (ls []T) {
	ls = make([]T, 0)
	for _, t := range x {
		if !fn(t) {
			continue
		}
		ls = append(ls, t)
	}
	return
}

func (x List[T]) FindOne(
	fn func(v T) bool,
) (res T, err error) {
	for _, t := range x {
		if fn(t) {
			res = t
		}
	}
	err = fnError.NewF(ErrNotFoundElement)
	return
}

func (x List[T]) FindIndex(fn func(v T) bool) int {
	for i, t := range x {
		if fn(t) {
			return i
		}
	}
	return -1
}

func (x List[T]) Has(v T) bool {
	for _, t := range x {
		if reflect.DeepEqual(t, v) {
			return true
		}
	}
	return false
}

func (x List[T]) HasAll(vs []T) bool {
	for _, v := range vs {
		if !x.Has(v) {
			return false
		}
	}
	return true
}
