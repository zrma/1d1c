package strings

import (
	"bufio"
	"encoding/csv"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/palindrome-index/problem", func() {
	type testData struct {
		s        string
		expected int32
	}

	DescribeTable("문제를 풀었다.",
		func(data testData) {

			actual := palindromeIndex(data.s)
			Expect(actual).Should(BeNumerically("==", data.expected))
		},
		Entry("aaab", testData{"aaab", 3}),
		Entry("baa", testData{"baa", 0}),
		Entry("aaa", testData{"aaa", -1}),
		Entry("a", testData{"a", -1}),
		Entry("abcadcdddaba", testData{"abcadcdddaba", -1}),
		Entry("abcd", testData{"abcd", -1}),
		Entry("quyjjdcgsvvsgcdjjyq", testData{"quyjjdcgsvvsgcdjjyq", 1}),
		Entry("hgygsvlfwcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcflvsgygh", testData{"hgygsvlfwcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcflvsgygh", 8}),
		Entry("fgnfnidynhxebxxxfmxixhsruldhsaobhlcggchboashdlurshxixmfxxxbexhnydinfngf", testData{"fgnfnidynhxebxxxfmxixhsruldhsaobhlcggchboashdlurshxixmfxxxbexhnydinfngf", 33}),
		Entry("bsyhvwfuesumsehmytqioswvpcbxyolapfywdxeacyuruybhbwxjmrrmjxwbhbyuruycaexdwyfpaloyxbcpwsoiqtymhesmuseufwvhysb", testData{"bsyhvwfuesumsehmytqioswvpcbxyolapfywdxeacyuruybhbwxjmrrmjxwbhbyuruycaexdwyfpaloyxbcpwsoiqtymhesmuseufwvhysb", 23}),
		Entry("fvyqxqxynewuebtcuqdwyetyqqisappmunmnldmkttkmdlnmnumppasiqyteywdquctbeuwenyxqxqyvf", testData{"fvyqxqxynewuebtcuqdwyetyqqisappmunmnldmkttkmdlnmnumppasiqyteywdquctbeuwenyxqxqyvf", 25}),
		Entry("mmbiefhflbeckaecprwfgmqlydfroxrblulpasumubqhhbvlqpixvvxipqlvbhqbumusaplulbrxorfdylqmgfwrpceakceblfhfeibmm", testData{"mmbiefhflbeckaecprwfgmqlydfroxrblulpasumubqhhbvlqpixvvxipqlvbhqbumusaplulbrxorfdylqmgfwrpceakceblfhfeibmm", 44}),
		Entry("tpqknkmbgasitnwqrqasvolmevkasccsakvemlosaqrqwntisagbmknkqpt", testData{"tpqknkmbgasitnwqrqasvolmevkasccsakvemlosaqrqwntisagbmknkqpt", 20}),
		Entry("lhrxvssvxrhl", testData{"lhrxvssvxrhl", -1}),
		Entry("prcoitfiptvcxrvoalqmfpnqyhrubxspplrftomfehbbhefmotfrlppsxburhyqnpfmqlaorxcvtpiftiocrp", testData{"prcoitfiptvcxrvoalqmfpnqyhrubxspplrftomfehbbhefmotfrlppsxburhyqnpfmqlaorxcvtpiftiocrp", 14}),
		Entry("kjowoemiduaaxasnqghxbxkiccikxbxhgqnsaxaaudimeowojk", testData{"kjowoemiduaaxasnqghxbxkiccikxbxhgqnsaxaaudimeowojk", -1}),
	)

	Context("equalPrefix 함수는", func() {
		Context("1글자 문자열 두 개가", func() {
			It("같으면 true를 반환한다.", func() {
				actual := equalPrefix("a", "a")
				Expect(actual).Should(BeTrue())
			})

			It("다르면 false를 반환한다.", func() {
				actual := equalPrefix("a", "b")
				Expect(actual).Should(BeFalse())
			})
		})

		Context("2글자 이상의 길이의 문자열 두 개의 앞 2글자", func() {
			It("같으면 true를 반환한다.", func() {

				actual := equalPrefix("abcd", "abab")
				Expect(actual).Should(BeTrue())
			})

			It("다르면 false를 반환한다.", func() {
				actual := equalPrefix("abc", "bbc")
				Expect(actual).Should(BeFalse())
			})
		})
	})

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long string", func() {

			file, err := os.Open("./test_data/palindrome_index.csv")
			Expect(err).ShouldNot(HaveOccurred())
			defer file.Close()

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			Expect(err).ShouldNot(HaveOccurred())

			var arr []string
			for _, row := range rows {
				arr = append(arr, row[0])
			}

			expected := []int32{
				16722,
				76661,
				10035,
				57089,
				46674,
			}

			Expect(arr).Should(HaveLen(5))
			for i, s := range arr {
				actual := palindromeIndex(s)
				Expect(actual).Should(BeNumerically("==", expected[i]))
			}
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 10), "시간 초과")
	}, 3)
})
