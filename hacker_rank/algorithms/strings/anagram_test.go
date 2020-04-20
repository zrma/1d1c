package strings

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/anagram/problem", func() {
	type testData struct {
		input    string
		expected int32
	}

	DescribeTable("문제를 풀었다.", func(d testData) {
		Expect(anagram(d.input)).Should(Equal(d.expected))
	},
		Entry("test_00", testData{"", -1}),
		Entry("test_01", testData{"aaabbb", 3}),
		Entry("test_02", testData{"ab", 1}),
		Entry("test_03", testData{"abc", -1}),
		Entry("test_04", testData{"mnop", 2}),
		Entry("test_05", testData{"xyyx", 0}),
		Entry("test_06", testData{"xaxbbbxx", 1}),
		Entry("test_07", testData{"hhpddlnnsjfoyxpciioigvjqzfbpllssuj", 10}),
		Entry("test_08", testData{"xulkowreuowzxgnhmiqekxhzistdocbnyozmnqthhpievvlj", 13}),
		Entry("test_09", testData{"dnqaurlplofnrtmh", 5}),
		Entry("test_10", testData{"aujteqimwfkjoqodgqaxbrkrwykpmuimqtgulojjwtukjiqrasqejbvfbixnchzsahpnyayutsgecwvcqngzoehrmeeqlgknnb", 26}),
		Entry("test_11", testData{"lbafwuoawkxydlfcbjjtxpzpchzrvbtievqbpedlqbktorypcjkzzkodrpvosqzxmpad", 15}),
		Entry("test_12", testData{"drngbjuuhmwqwxrinxccsqxkpwygwcdbtriwaesjsobrntzaqbe", -1}),
		Entry("test_13", testData{"ubulzt", 3}),
		Entry("test_14", testData{"vxxzsqjqsnibgydzlyynqcrayvwjurfsqfrivayopgrxewwruvemzy", 13}),
		Entry("test_15", testData{"xtnipeqhxvafqaggqoanvwkmthtfirwhmjrbphlmeluvoa", 13}),
		Entry("test_16", testData{"gqdvlchavotcykafyjzbbgmnlajiqlnwctrnvznspiwquxxsiwuldizqkkaawpyyisnftdzklwagv", -1}),
	)
})
