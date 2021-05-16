package lv1medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoI(t *testing.T) {
	t.Log("https://leetcode.com/problems/string-to-integer-atoi/")

	for _, tt := range []struct {
		description string
		given       string
		want        int
	}{
		{"정수", "42", 42},
		{"공백 + 음수", "   -42", -42},
		{"정수 + 문자", "4193 with words", 4193},
		{"문자 + 정수", "words and 987", 0},
		{"음수", "-91283472332", -2147483648},
		{"없음", "", 0},
		{"소수", "3.14159", 3},
		{"공백 + 음수 + 문자", "  -0012a42", -12},
		{"공백", "          ", 0},
		{"공백 + 기호", "          -", 0},
		{"+ 명시 양수", "+1", 1},
		{"큰 수 overflow", "9223372036854775808", 2147483647},
		{"큰 수 overflow", "18446744073709551617", 2147483647},
		{"앞의 공백 + 0 버림", "  0000000000012345678", 12345678},
		{"0으로 채워진 음수", "-000000000000001", -1},
	} {
		t.Run(tt.description, func(t *testing.T) {
			got := myAtoI(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
