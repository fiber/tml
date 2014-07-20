package tml

// StrLen returns the length of a string as a number of unicode characters
// rather than number of bytes that len(s) does
func StrLen(s string) int {
	i := 0
	for _ = range s {
		i++
	}
  return i
}
