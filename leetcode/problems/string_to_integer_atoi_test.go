package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/string-to-integer-atoi/", func() {
	type testData struct {
		input    string
		expected int
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		actual := myAtoI(data.input)
		Expect(actual).Should(Equal(data.expected))
	},
		Entry("정수", testData{"42", 42}),
		Entry("공백 + 음수", testData{"   -42", -42}),
		Entry("정수 + 문자", testData{"4193 with words", 4193}),
		Entry("문자 + 정수", testData{"words and 987", 0}),
		Entry("음수", testData{"-91283472332", -2147483648}),
		Entry("없음", testData{"", 0}),
		Entry("소수", testData{"3.14159", 3}),
		Entry("공백 + 음수 + 문자", testData{"  -0012a42", -12}),
		Entry("공백", testData{"          ", 0}),
		Entry("공백 + 기호", testData{"          -", 0}),
		Entry("+ 명시 양수", testData{"+1", 1}),
		Entry("큰 수 overflow", testData{"9223372036854775808", 2147483647}),
		Entry("큰 수 overflow", testData{"18446744073709551617", 2147483647}),
		Entry("앞의 공백 + 0 버림", testData{"  0000000000012345678", 12345678}),
		Entry("0으로 채워진 음수", testData{"-000000000000001", -1}),
	)
})
