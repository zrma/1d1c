package lv_0_easy

func longestCommonPrefix(stringArray []string) string {
	if len(stringArray) == 0 {
		return ""
	}

	first := stringArray[0]
	pos := func() int {
		var cur int
		for {
			if cur == len(first) {
				return cur
			}
			c := first[cur]
			for i := 1; i < len(stringArray); i++ {
				s := stringArray[i]
				if len(s) == cur || s[cur] != c {
					return cur
				}
			}
			cur++
		}
	}()
	if pos == 0 {
		return ""
	}

	return first[0:pos]
}
