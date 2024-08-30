package main

import (
	"day_4/utils"
	"fmt"
)


func main() {
    secretKey := "ckczppom"
	secretKey2 := "ckczppom"
    prefix := "00000"
	prefix2 := "000000"

    result := utils.MineAdventCoins(secretKey, prefix)
	result2 := utils.MineAdventCoins(secretKey2, prefix2)

    fmt.Printf("The lowest number is: %d\n", result)
	fmt.Printf("The lowest number is: %d\n", result2)
}
