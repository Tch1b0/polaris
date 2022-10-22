package general

import "os"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// read and process the input
func ProcessInput[T any](path string, processor func(string) T) (T, error) {
	file, err := os.Open(path)
	var content []byte
	file.Read(content)

	return processor(string(content)), err
}
