package main

import (
	"day_11/utils"
	"fmt"
)

func main() {
	
	currentPassword := "cqjxjnds"
	currentPassword2 := "cqjxxyzz"
	
	nextPassword := utils.FindNextPassword(currentPassword)
	nextPassword2 := utils.FindNextPassword(currentPassword2)

	fmt.Printf("Santa's next password is: %s\n", nextPassword)
	fmt.Printf("Santa's next password is: %s\n", nextPassword2)
}
