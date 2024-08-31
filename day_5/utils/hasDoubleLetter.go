package utils

func HasDoubleLetter(s string) bool {
    for i := 0; i < len(s)-1; i++ {
        if s[i] == s[i+1] {
            return true
        }
    }
    return false
}
