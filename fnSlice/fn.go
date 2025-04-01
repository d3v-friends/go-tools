package fnSlice

import (
	"github.com/d3v-friends/go-tools/fnError"
	"math/rand/v2"
)

const (
	ErrNotFoundElement = "not_found_element"
)

func Has[T any](
	vs []T,
	fn func(v T) bool,
) bool {
	for _, v := range vs {
		if fn(v) {
			return true
		}
	}
	return false
}

func Filter[T any](
	vs []T,
	fn func(v T) bool,
) (ls []T) {
	ls = make([]T, 0)
	for _, v := range vs {
		if !fn(v) {
			continue
		}
		ls = append(ls, v)
	}
	return
}

func Find[T any](
	vs []T,
	fn func(v T) bool,
) (ls []T) {
	ls = make([]T, 0)
	for _, v := range vs {
		if !fn(v) {
			continue
		}
		ls = append(ls, v)
	}
	return
}

func FindOne[T any](
	vs []T,
	fn func(v T) bool,
) (res T, err error) {
	for _, v := range vs {
		if fn(v) {
			res = v
			return
		}
	}
	err = fnError.New(ErrNotFoundElement)
	return
}

func Concat[T any](
	vs ...[]T,
) (res []T) {
	res = make([]T, 0)

	for _, v := range vs {
		res = append(res, v...)
	}

	return
}

func Divide[T any](
	vs []T,
	unit int,
) (res [][]T) {
	res = make([][]T, 0)

	var page = len(vs) / unit
	if len(vs)%unit != 0 {
		page += 1
	}

	for i := 0; i < page; i++ {
		if i == page-1 {
			res = append(res, vs[i*unit:])
		} else {
			res = append(res, vs[i*unit:(i+1)*unit])
		}
	}

	return
}

func Parse[A, B any](
	vs []A,
	parser func(v A) (B, error),
) (ls []B, err error) {
	ls = make([]B, len(vs))
	for i, v := range vs {
		if ls[i], err = parser(v); err != nil {
			return
		}
	}
	return
}

func Each[A, B any](
	vs []A,
	parser func(v A) B,
) (ls []B) {
	ls = make([]B, len(vs))
	for i, v := range vs {
		ls[i] = parser(v)
	}
	return
}

func Deduplicate[T any](
	vs []T,
	isSame func(a T, b T) bool,
) (res []T) {
	res = make([]T, 0)
	for _, in := range vs {
		var has = false

		for _, out := range res {
			if isSame(in, out) {
				has = true
				break
			}
		}

		if !has {
			res = append(res, in)
		}
	}

	return
}

func Map[K comparable, V, R any](
	m map[K]V,
	fn func(K, V) R,
) (ls []R) {
	ls = make([]R, len(m))
	var i = 0
	for k, v := range m {
		ls[i] = fn(k, v)
	}
	return
}

func ShuffleKnuth[T any](
	vs []T,
) (res []T) {
	res = make([]T, len(vs))
	for i, v := range vs {
		res[i] = v
	}

	var total = len(vs)
	for i := 0; i < total; i++ {
		var c = res[i]
		var nextIdx = i + rand.IntN(total-i)
		res[i] = res[nextIdx]
		res[nextIdx] = c
	}

	return
}
