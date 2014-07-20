package tml

// StrLen returns the number of unicode characters in a string
// rather than number of bytes that len(s) does
func StrLen(s string) int {
	i := 0
	for _ = range s {
		i++
	}
	return i
}
