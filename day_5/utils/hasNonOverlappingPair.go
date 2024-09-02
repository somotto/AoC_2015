package utils

func HasNonOverlappingPair(s string) bool {
    n := len(s)
    for i := 0; i < n-1; i++ {
        pair := s[i:i+2]
        for j := i + 2; j < n-1; j++ {
            if s[j:j+2] == pair {
                return true
            }
        }
    }
    return false
}
