package houses

import "fmt"

func CountHousesWithSantaRobo(directions string) int {
	housesVisited := make(map[string]bool)

	xS, yS := 0, 0
	xR, yR := 0, 0

	housesVisited["0,0"] = true

	for i, dir := range directions {
		// Santa's turn
		if i%2 == 0 {
			switch dir {
			case '^':
				yS++
			case 'v':
				yS--
			case '>':
				xS++
			case '<':
				xS--
			}
			housesVisited[fmt.Sprintf("%d,%d", xS, yS)] = true
			// Robo-Santa's turn
		} else {
			switch dir {
			case '^':
				yR++
			case 'v':
				yR--
			case '>':
				xR++
			case '<':
				xR--
			}
			housesVisited[fmt.Sprintf("%d,%d", xR, yR)] = true
		}
	}
	return len(housesVisited)
}
