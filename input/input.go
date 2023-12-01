package input

import "os"

// read and process the input
func Process[T any](path string, processor func(string) T) T {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return processor(string(content))
}

func Read(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func Memoize[T comparable, N any](fn func(T) N) func(T) N {
	cache := make(map[T]N)
	return func(in T) N {
		if val, ok := cache[in]; ok {
			return val
		}

		cache[in] = fn(in)
		return cache[in]
	}
}
