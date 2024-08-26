package Not_quite_lisp

func Position(input string) int {
	floor := 0

	for i, step := range input {
		if step == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}
