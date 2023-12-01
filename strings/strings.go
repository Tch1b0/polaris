package strings

import (
	"strconv"

	"github.com/Tch1b0/polaris/math"
)

func Reverse(str string) string {
	revStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}

	return revStr
}

// returns a sorted version of the string
func Sort(str string) string {
	return string(math.Sort([]byte(str)))
}

// returns whether the strings are anagrams of each other
func IsAnagram(str1, str2 string) bool {
	return Sort(str1) == Sort(str2)
}

func Contains(str string, sub string) bool {
	for i := 0; i <= len(str)-len(sub)-1; i++ {
		block := str[i : i+len(sub)]
		if block == sub {
			return true
		}
	}

	return false
}

func Itoa(n int) string {
	return strconv.Itoa(n)
}

func Atoi(str string) (int, error) {
	return strconv.Atoi(str)
}

func ToRunes(str string) []rune {
	runes := make([]rune, len(str))
	for i, r := range str {
		runes[i] = r
	}

	return runes
}
