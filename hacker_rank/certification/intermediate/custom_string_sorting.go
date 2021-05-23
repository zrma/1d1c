package intermediate

import (
	"sort"
	"strings"
)

func customSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		a := strArr[i]
		b := strArr[j]

		isAEven := len(a)%2 == 0
		isBEven := len(b)%2 == 0

		if len(a) == len(b) {
			if isBEven {
				return strings.Compare(a, b) < 0
			} else {
				return strings.Compare(a, b) > 0
			}
		}

		switch {
		case isAEven && isBEven:
			return len(a) < len(b)
		case isAEven && !isBEven:
			return true
		case !isAEven && isBEven:
			return false
		default:
			return len(a) > len(b)
		}
	})

	return strArr
}
