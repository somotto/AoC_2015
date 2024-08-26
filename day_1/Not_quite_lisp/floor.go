package Not_quite_lisp

func Floor(input string) int {

	floor := 0

	for _, step := range input {
		if step == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}
