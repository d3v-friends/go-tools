package fnSlice

type Number interface {
	int | int8 | int16 | int32 | int64
}

func Page[N Number](total, size N) (page N) {
	page = total / size
	if total%size != 0 {
		page += 1
	}
	return
}
