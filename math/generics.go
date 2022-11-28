package math

import (
	"fmt"

	"github.com/Tch1b0/polaris/array"
)

func Min[T Number](arr []T) T {
	if len(arr) == 0 {
		panic("Invalid array length")
	}

	m := arr[0]

	for _, value := range arr {
		if value < m {
			m = value
		}
	}

	return m
}

func Max[T Number](arr []T) T {
	if len(arr) == 0 {
		panic("Invalid array length")
	}

	m := arr[0]

	for _, value := range arr {
		if value > m {
			m = value
		}
	}

	return m
}

func Sum[T Number](arr []T) T {
	sum := T(0)

	for _, value := range arr {
		sum += value
	}

	return sum
}

func Product[T Number](arr []T) T {
	prod := T(1)

	for _, value := range arr {
		prod *= value
	}

	return prod
}

func Average[T Number](arr []T) T {
	sum := Sum(arr)
	return sum / T(len(arr))
}

func Median[T Number](arr []T) T {
	return arr[int(len(arr)/2)]
}

// sort array using an insertion sort algorithm
func Sort[T Number](arr []T) []T {
	for i, value := range arr {
		for j := 0; j < i; j++ {
			if arr[j] > value {
				arr = array.Move(arr, i, j)
				break
			}
		}
	}

	return arr
}
