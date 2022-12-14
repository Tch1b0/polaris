package array

func Map[T any, N any](arr []T, predicate func(T, int) N) []N {
	newArr := make([]N, len(arr))
	for i, item := range arr {
		newArr[i] = predicate(item, i)
	}

	return newArr
}

func Filter[T any](arr []T, predicate func(T, int) bool) []T {
	newArr := make([]T, len(arr))
	for i, item := range arr {
		if predicate(item, i) {
			newArr = append(newArr, item)
		}
	}

	return newArr
}

func Some[T any](arr []T, predicate func(T, int) bool) bool {
	for i, item := range arr {
		if predicate(item, i) {
			return true
		}
	}

	return false
}

func All[T any](arr []T, predicate func(T, int) bool) bool {
	for i, item := range arr {
		if !predicate(item, i) {
			return false
		}
	}

	return true
}

func GetFirst[T any](arr []T, predicate func(T, int) bool) *T {
	for i, item := range arr {
		if predicate(item, i) {
			return &item
		}
	}

	return nil
}

func GetAll[T any](arr []T, predicate func(T, int) bool) []T {
	matches := []T{}

	for i, item := range arr {
		if predicate(item, i) {
			matches = append(matches, item)
		}
	}

	return matches
}

func None[T any](arr []T, predicate func(T, int) bool) bool {
	for i, item := range arr {
		if predicate(item, i) {
			return false
		}
	}

	return true
}

func Each[T any](arr []T, predicate func(T, int)) {
	for i, item := range arr {
		predicate(item, i)
	}
}
