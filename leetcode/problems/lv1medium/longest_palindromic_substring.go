package lv1medium

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

type checkOptions struct {
	index  int
	offset int

	maxLength int
	min, max  int

	length int
	s      string
}

func checkByRange(opts *checkOptions) {
	start := opts.index - opts.offset
	if start < 0 {
		return
	}

	for i := -1; i < 2; i++ {
		end := opts.index + opts.offset + i
		if end >= opts.length {
			break
		}

		if isPalindrome(opts.s, start, end) && (end-start+1) > opts.maxLength {
			opts.maxLength = end - start + 1
			opts.min = start
			opts.max = end
		}
	}
}

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 || (length <= 3 && s[0] == s[length-1]) {
		return s
	}

	opts := &checkOptions{
		maxLength: 1,
		length:    length,
		s:         s,
	}
	for opts.index = 1; opts.index < length-1; opts.index++ {
		for opts.offset = 0; opts.offset <= length/2; opts.offset++ {
			checkByRange(opts)
		}
	}
	return s[opts.min : opts.max+1]
}
