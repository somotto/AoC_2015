package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5 to calculate the MD5 hash of a string
func Md5Hash(data string) [16]byte {
    return md5.Sum([]byte(data))
}

// ToHexString to convert a byte slice to a hexadecimal string
func ToHexString(data []byte) string {
    return fmt.Sprintf("%x", data)
}

// HashHasPrefix check if the hash starts with the required prefix
func HashHasPrefix(hash string, prefix string, prefixLen int) bool {
    if len(hash) < prefixLen {
        return false
    }

    for i := 0; i < prefixLen; i++ {
        if hash[i] != prefix[i] {
            return false
        }
    }
    return true
}