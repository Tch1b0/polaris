package maps

func HasKey[T comparable, N any](m map[T]N, searched T) bool {
	for k := range m {
		if k == searched {
			return true
		}
	}

	return false
}

func HasValue[T comparable, N comparable](m map[T]N, searched N) bool {
	for _, v := range m {
		if v == searched {
			return true
		}
	}

	return false
}

func Keys[T comparable, N any](m map[T]N) []T {
	keys := make([]T, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func Values[T comparable, N any](m map[T]N) []N {
	values := make([]N, len(m))
	for _, v := range m {
		values = append(values, v)
	}

	return values
}
