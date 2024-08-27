package utils

func SmallestSideArea(l, w, h int) int {
	side1 := l * w
	side2 := w * h
	side3 := h * l

	smallest := side1

	if side2 < smallest {
		smallest = side2
	}
	if side3 < smallest {
		smallest = side3
	}
	return smallest
}
