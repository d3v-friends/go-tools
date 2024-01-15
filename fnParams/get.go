package fnParams

func Get[T any](v []T, defs ...T) T {
	if len(v) == 0 {
		if len(defs) == 0 {
			return *new(T)
		}
		return defs[0]
	}
	return v[0]
}
