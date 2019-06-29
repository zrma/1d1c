package strings

import (
	"fmt"
	"strings"
)

func superReducedString(s string) string {
	const emptyString = "Empty String"

	if len(s) == 0 {
		return emptyString
	}

	pair := make(map[string]interface{})
	for _, c := range s {
		pair[fmt.Sprintf("%c%c", c, c)] = nil
	}

	find := func() (string, bool) {
		for c := range pair {
			if strings.Contains(s, c) {
				return c, true
			}
		}
		return "", false
	}

	for c, ok := find(); ok; c, ok = find() {
		s = strings.ReplaceAll(s, c, "")
	}

	if s == "" {
		return emptyString
	}
	return s
}
