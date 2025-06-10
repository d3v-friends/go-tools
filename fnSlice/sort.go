package fnSlice

// Sort
// 복제하여 정렬한다. selection sort -> chatGPT 에서 작성
// a > b 이면 올림차순
// a < b 이면 내림차순
func Sort[T any](ls []T, fn func(a, b T) bool) (res []T) {
	res = make([]T, len(ls))
	copy(res, ls)

	if len(res) == 1 {
		return
	}

	for i := 0; i < len(res)-1; i++ {
		var minIdx = i

		for j := i + 1; j < len(res); j++ {
			if fn(res[minIdx], res[j]) {
				minIdx = j
			}
		}

		if minIdx != i {
			res[i], res[minIdx] = res[minIdx], res[i]
		}
	}

	return
}
