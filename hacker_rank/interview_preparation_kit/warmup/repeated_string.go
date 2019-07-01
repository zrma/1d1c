package warmup

import "strings"

const findTarget = "a"

func repeatedString(s string, n int64) int64 {
	if !isValid(s, n) {
		return 0
	}

	length := int64(len(s))
	cnt := int64(strings.Count(s, findTarget))
	repeatCnt := n / length
	mod := n % length

	return cnt*repeatCnt + partialCnt(s, mod)
}

func partialCnt(s string, n int64) int64 {
	if !isValid(s, n) {
		return 0
	}

	return int64(strings.Count(s[:n], findTarget))
}

func isValid(s string, n int64) bool {
	if n == 0 {
		return false
	}

	if s == "" {
		return false
	}

	return true
}
