package lv1medium

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	buffers := make([][]rune, numRows)

	var downward bool
	var pos int
	for _, c := range s {
		if pos == 0 {
			downward = true
		}
		if pos == numRows-1 {
			downward = false
		}

		buffers[pos] = append(buffers[pos], c)

		if downward {
			pos++
		} else {
			pos--
		}
	}

	var result string
	for _, buffer := range buffers {
		result += string(buffer)
	}

	return result
}
