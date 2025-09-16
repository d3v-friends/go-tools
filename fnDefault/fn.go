package fnDefault

func Param[T any](ls []T, value T) T {
	if len(ls) == 1 {
		return ls[0]
	}
	return value
}
