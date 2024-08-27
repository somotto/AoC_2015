package utils

func WrappingPaper(l, w, h int) int {
	surfaceArea := 2 * (l*w + w*h + h*l)
	smallestSide := SmallestSideArea(l, w, h)
	return surfaceArea + smallestSide
}
