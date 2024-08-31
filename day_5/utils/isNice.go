package utils

func IsNice(s string) bool {
    if ContainsDisallowedSubstring(s) {
        return false
    }
    if CountVowels(s) < 3 {
        return false
    }
    if !HasDoubleLetter(s) {
        return false
    }
    return true
}
