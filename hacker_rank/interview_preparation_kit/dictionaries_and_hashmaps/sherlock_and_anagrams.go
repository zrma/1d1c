package dictionaries_and_hashmaps

import (
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

func sherlockAndAnagrams(s string) int32 {
	length := len(s)

	m := make(map[string]int32)
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			token := s[i:j]
			token = utils.SortString(token)
			m[token]++
		}
	}

	var total int32
	for _, v := range m {
		// v Combinations of 2
		// vC2 = v * (v - 1) / (2 * 1)
		total += v * (v - 1) / 2
	}
	return total
}
