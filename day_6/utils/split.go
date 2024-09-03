package utils

// Split splits a string into substrings separated by a delimiter.
func Split(s, sep string) []string {
	var substrings []string
	start := 0
	sepLen := len(sep)

	for i := 0; i <= len(s)-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			substrings = append(substrings, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}
	substrings = append(substrings, s[start:])
	return substrings
}