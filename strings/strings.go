package strings

func Reverse(str string) string {
	revStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}

	return revStr
}
