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

func Invert[T comparable, N comparable](m map[T]N) map[N]T {
	n := make(map[N]T, len(m))
	for k, v := range m {
		n[v] = k
	}

	return n
}

func FromArray[T comparable, N any](keys []T, values []N) map[T]N {
	m := make(map[T]N, len(keys))

	for i, key := range keys {
		m[key] = values[i]
	}

	return m
}

func ToArray[T comparable, N any](m map[T]N) ([]T, []N) {
	keys := []T{}
	vals := []N{}

	for key, value := range m {
		keys = append(keys, key)
		vals = append(vals, value)
	}

	return keys, vals
}
