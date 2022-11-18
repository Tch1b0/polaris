package input_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Tch1b0/polaris/input"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("abc")

	require.Equal(t, input.Read(tmpFile.Name()), "abc")
}

func TestProcessing(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("1 2 3")

	require.Equal(
		t,
		input.Process(tmpFile.Name(), func(in string) []int {
			res := make([]int, 3)
			for i, char := range strings.Split(in, " ") {
				res[i], _ = strconv.Atoi(char)
			}

			return res
		}),
		[]int{1, 2, 3},
	)
}
