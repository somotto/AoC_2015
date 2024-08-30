package utils

import (
	"fmt"
)

// Function to find the lowest number that produces a hash with the given prefix
func MineAdventCoins(secretKey string, prefix string) int {
    i := 1
    prefixLen := len(prefix)
    for {
        input := fmt.Sprintf("%s%d", secretKey, i)
        hash := Md5Hash(input)
        hashHex := ToHexString(hash[:])

        if HashHasPrefix(hashHex, prefix, prefixLen) {
            return i
        }
        i++
    }
}