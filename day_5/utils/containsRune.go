package utils

// ContainsRune checks if the rune 'r' is in the string 's'.
func ContainsRune(s string, r rune) bool {
    for _, ch := range s {
        if ch == r {
            return true
        }
    }
    return false
}