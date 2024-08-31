package utils

// containsSubstring checks if the substring 'substr' is in the string 's'.
func ContainsSubstring(s, substr string) bool {
    if len(substr) == 0 || len(substr) > len(s) {
        return false
    }
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return true
        }
    }
    return false
}