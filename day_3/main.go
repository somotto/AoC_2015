package main

import (
	"day_3/houses"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		return
	}

	data, err := os.ReadFile("direction.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	directions := string(data)
	housesRecievedMore := houses.CountHouses(directions)

	housesRecievedMoreWithSantaandRobo := houses.CountHousesWithSantaRobo(directions)

	fmt.Println(housesRecievedMore)

	fmt.Println(housesRecievedMoreWithSantaandRobo)

}
