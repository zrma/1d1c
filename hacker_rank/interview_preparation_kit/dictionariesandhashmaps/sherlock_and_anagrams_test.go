package dictionariesandhashmaps

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem", func() {
	DescribeTable("문제를 풀었다",
		func(s string, i int) {
			actual := sherlockAndAnagrams(s)
			Expect(actual).Should(BeNumerically("==", i))
		},
		Entry("test_0", "abba", 4),
		Entry("test_1", "abcd", 0),
		Entry("test_2", "ifailuhkqq", 3),
		Entry("test_3", "kkkk", 10),
		Entry("test_4", "cdcd", 5),
	)

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long string", func() {
			for _, data := range []struct {
				s        string
				expected int
			}{
				{"dbcfibibcheigfccacfegicigcefieeeeegcghggdheichgafhdigffgifidfbeaccadabecbdcgieaffbigffcecahafcafhcdg", 1464},
				{"dfcaabeaeeabfffcdbbfaffadcacdeeabcadabfdefcfcbbacadaeafcfceeedacbafdebbffcecdbfebdbfdbdecbfbadddbcec", 2452},
				{"gjjkaaakklheghidclhaaeggllagkmblhdlmihmgkjhkkfcjaekckaafgabfclmgahmdebjekaedhaiikdjmfbmfbdlcafamjbfe", 873},
				{"fdbdidhaiqbggqkhdmqhmemgljaphocpaacdohnokfqmlpmiioedpnjhphmjjnjlpihmpodgkmookedkehfaceklbljcjglncfal", 585},
				{"bcgdehhbcefeeadchgaheddegbiihehcbbdffiiiifgibhfbchffcaiabbbcceabehhiffggghbafabbaaebgediafabafdicdhg", 1305},
				{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 166650},
				{"mhmgmbbccbbaffhbncgndbffkjbhmkfncmihhdhcebmchnfacdigflhhbekhfejblegakjjiejeenibemfmkfjbkkmlichlkbnhc", 840},
				{"fdacbaeacbdbaaacafdfbbdcefadgfcagdfcgbgeafbfbggdedfbdefdbgbefcgdababafgffedbefdecbaabdaafggceffbacgb", 2134},
				{"bahdcafcdadbdgagdddcidaaicggcfdbfeeeghiibbdhabdhffddhffcdccfdddhgiceciffhgdibfdacbidgagdadhdceibbbcc", 1571},
				{"dichcagakdajjhhdhegiifiiggjebejejciaabbifkcbdeigajhgfcfdgekfajbcdifikafkgjjjfefkdbeicgiccgkjheeiefje", 1042},
			} {
				actual := sherlockAndAnagrams(data.s)
				Expect(actual).Should(BeNumerically("==", data.expected))
			}
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 10), "시간 초과")
	}, 3)
})
