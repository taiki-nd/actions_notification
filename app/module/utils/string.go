package utils

// GetPrefix returns the first n characters of a string.
func GetPrefix(s string, num int) string {
	if len(s) < num {
		return s
	}
	return s[:num]
}
