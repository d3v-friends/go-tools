package fnSlice

import "math/rand"

func PickRand[T any](ls []T, size int) (res []T) {
	if len(ls) < size {
		return ls
	}

	var has = make(map[int]bool)

	for {
		var idx = rand.Intn(len(ls))
		var _, exist = has[idx]
		if exist {
			continue
		}

		has[idx] = true
		res = append(res, ls[idx])

		if len(res) == size {
			break
		}
	}

	return
}
