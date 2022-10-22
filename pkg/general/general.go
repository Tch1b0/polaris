package general

import "os"

func ReadInput(path string) (string, error) {
	file, err := os.Open(path)
	var content []byte
	file.Read(content)

	return string(content), err
}
