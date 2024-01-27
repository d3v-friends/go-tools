package fnReflect

func Pointer[T any](v T) (res *T) {
	return &v
}
