package general

import "os"

// read and process the input
func ProcessInput[T any](path string, processor func(string) T) (T, error) {
	file, err := os.Open(path)
	var content []byte
	file.Read(content)

	return processor(string(content)), err
}
