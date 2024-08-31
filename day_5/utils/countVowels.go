package utils

func CountVowels(s string) int {
    vowels := "aeiou"
    count := 0
    for _, ch := range s {
        if ContainsRune(vowels, ch) {
            count++
        }
    }
    return count
}
