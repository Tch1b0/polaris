package math

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
