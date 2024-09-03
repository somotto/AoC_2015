package utils

// Fields splits a string into fields based on whitespace without using boolean flags.
func Fields(s string) []string {
	var fields []string
	var field string

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			// If we encounter whitespace and field is not empty, add it to the fields list
			if len(field) > 0 {
				fields = append(fields, field)
				field = ""
			}
		} else {
			// Append non-whitespace characters to the current field
			field += string(ch)
		}
	}

	// If there's any field left after the loop, add it to the fields list
	if len(field) > 0 {
		fields = append(fields, field)
	}

	return fields
}
