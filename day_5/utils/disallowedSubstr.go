package utils

func ContainsDisallowedSubstring(s string) bool {
    disallowed := []string{"ab", "cd", "pq", "xy"}
    
    for _, substr := range disallowed {
        if ContainsSubstring(s, substr) {
            return true
        }
    }
    return false
}


