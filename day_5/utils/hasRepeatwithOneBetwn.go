package utils

func HasRepeatWithOneBetween(s string) bool {
    for i := 0; i < len(s)-2; i++ {
        if s[i] == s[i+2] {
            return true
        }
    }
    return false
}
