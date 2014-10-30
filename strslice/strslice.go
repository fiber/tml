package strslice

// Contains returns true is list contains s
func Contains(list []string, s string) bool {
	for _, n := range list {
		if n == s {
			return true
		}
	}
	return false
}

