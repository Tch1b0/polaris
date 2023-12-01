package runes

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsAlpha(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func IsAlNum(r rune) bool {
	return IsDigit(r) || IsAlpha(r)
}

func DigitToI(r rune) int {
	return int(r - '0')
}
