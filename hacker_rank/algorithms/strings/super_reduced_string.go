package strings

func superReducedString(s string) string {
	const emptyString = "Empty String"

	if len(s) == 0 {
		return emptyString
	}

	stack := make([]rune, 0)
	idx := -1

	for _, r := range s {
		if idx >= 0 && stack[idx] == r {
			stack = stack[0:idx]
			idx--
		} else {
			stack = append(stack, r)
			idx++
		}
	}

	if len(stack) == 0 {
		return emptyString
	}
	return string(stack)
}
