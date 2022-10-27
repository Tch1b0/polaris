package input

import (
	"io/ioutil"
)

// read and process the input
func Process[T any](path string, processor func(string) T) T {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return processor(string(content))
}

func Read(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}
