package dictionariesandhashmaps

import (
	"github.com/zrma/going/utils/str"
)

func twoStrings(s1, s2 string) string {
	s1 = str.Sort(s1)
	s2 = str.Sort(s2)

	var i, j int
	for i < len(s1)-1 && j < len(s2)-1 {
		diff := int(s1[i]) - int(s2[j])

		if diff < 0 {
			i++
		} else if diff > 0 {
			j++
		} else {
			return "YES"
		}
	}

	return "NO"
}
