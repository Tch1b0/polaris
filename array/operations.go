package array

// returns the position of value in arr or -1 if it does not exist
func Index[T comparable](arr []T, value T) int {
	for i, item := range arr {
		if item == value {
			return i
		}
	}

	return -1
}

func Reverse[T any](arr []T) []T {
	arrLen := len(arr)
	n := make([]T, arrLen)

	for i, item := range arr {
		n[arrLen-i-1] = item
	}

	return n
}

func Remove[T comparable](arr []T, target T) []T {
	n := []T{}

	for _, item := range arr {
		if item != target {
			n = append(n, item)
		}
	}

	return n
}

func Pop[T any](arr []T, idx int) []T {
	n := []T{}

	for i, item := range arr {
		if i != idx {
			n = append(n, item)
		}
	}

	return n
}

// Mutation Pop: mutates the given array and returns the removed item
func MPop[T any](arr *[]T, idx int) T {
	arrLen := len(*arr)
	item := (*arr)[idx]

	for i := idx; i < arrLen-1; i++ {
		(*arr)[i] = (*arr)[i+1]
	}

	(*arr) = (*arr)[:arrLen-1]

	return item
}

func Insert[T any](arr []T, newItem T, idx int) []T {
	n := []T{}
	for i, value := range arr {
		if i == idx {
			n = append(n, newItem)
		}

		n = append(n, value)
	}

	return n
}

func Equals[T comparable](arr1 []T, arr2 []T) bool {
	for i, value := range arr1 {
		if value != arr2[i] {
			return false
		}
	}

	return true
}

func Move[T any](arr []T, from int, to int) []T {
	n := []T{}
	moveVal := arr[from]

	for i, value := range arr {
		if i == from {
			continue
		} else if i == to {
			n = append(n, moveVal)
		}

		n = append(n, value)
	}

	return n
}
