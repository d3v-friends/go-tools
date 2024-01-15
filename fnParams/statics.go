package fnParams

func Between(min, max, value int) bool {
	if value < min {
		return false
	}

	if max < value {
		return false
	}

	return true
}
