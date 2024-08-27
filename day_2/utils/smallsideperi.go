package utils

func FindRibbonFeet(l, w, h int) int {
	side1 := 2 * (l + w)
	side2 := 2 * (w + h)
	side3 := 2 * (h + l)

	smallestSide := side1

	if side2 < smallestSide {
		smallestSide = side2
	}
	if side3 < smallestSide {
		smallestSide = side3
	}
	bowlength := l*w*h
	return smallestSide + bowlength
}
