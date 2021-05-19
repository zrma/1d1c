package strings

import (
	"1d1go/utils/str"
)

func anagram(s string) int32 {
	length := len(s)
	if length < 1 || length%2 == 1 {
		return -1
	}

	s1 := str.Sort(s[0 : length/2])
	s2 := str.Sort(s[length/2:])

	length = length / 2
	var eq int32
	var i, j int
	for i < length && j < length {
		diff := int(s1[i]) - int(s2[j])

		if diff < 0 {
			i++
		} else if diff > 0 {
			j++
		} else {
			eq++
			i++
			j++
		}
	}

	return int32(length) - eq
}
