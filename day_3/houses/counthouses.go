package houses

import "fmt"

func CountHouses(directions string) int {
	housesVisited := make(map[string]bool)

	x, y := 0, 0

	//mark the starting house as visited
	housesVisited[fmt.Sprintf("%d,%d", x, y)] = true

	//Process each direction command
	for _, direction := range directions {
		switch direction {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}
		housesVisited[fmt.Sprintf("%d,%d", x, y)] = true
	}
	return len(housesVisited)
}
