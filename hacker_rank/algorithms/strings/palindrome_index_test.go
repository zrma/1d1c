package strings

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeIndex(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/palindrome-index/problem")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want int32
	}{
		{"aaab", 3},
		{"baa", 0},
		{"aaa", -1},
		{"a", -1},
		{"abcadcdddaba", -1},
		{"abcd", -1},
		{"quyjjdcgsvvsgcdjjyq", 1},
		{"hgygsvlfwcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcflvsgygh", 8},
		{"fgnfnidynhxebxxxfmxixhsruldhsaobhlcggchboashdlurshxixmfxxxbexhnydinfngf", 33},
		{"bsyhvwfuesumsehmytqioswvpcbxyolapfywdxeacyuruybhbwxjmrrmjxwbhbyuruycaexdwyfpaloyxbcpwsoiqtymhesmuseufwvhysb", 23},
		{"fvyqxqxynewuebtcuqdwyetyqqisappmunmnldmkttkmdlnmnumppasiqyteywdquctbeuwenyxqxqyvf", 25},
		{"mmbiefhflbeckaecprwfgmqlydfroxrblulpasumubqhhbvlqpixvvxipqlvbhqbumusaplulbrxorfdylqmgfwrpceakceblfhfeibmm", 44},
		{"tpqknkmbgasitnwqrqasvolmevkasccsakvemlosaqrqwntisagbmknkqpt", 20},
		{"lhrxvssvxrhl", -1},
		{"prcoitfiptvcxrvoalqmfpnqyhrubxspplrftomfehbbhefmotfrlppsxburhyqnpfmqlaorxcvtpiftiocrp", 14},
		{"kjowoemiduaaxasnqghxbxkiccikxbxhgqnsaxaaudimeowojk", -1},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := palindromeIndex(tt.s)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestEqualPrefix(t *testing.T) {
	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		description    string
		given1, given2 string
		want           bool
	}{
		{
			description: "한 글자 문자열, 같으면 true",
			given1:      "a", given2: "a",
			want: true,
		},
		{
			description: "한 글자 문자열, 다르면 false",
			given1:      "a", given2: "b",
			want: false,
		},
		{
			description: "두 글자 이상, 앞 2글자 같으면 true",
			given1:      "abcd", given2: "abab",
			want: true,
		},
		{
			description: "두 글자 이상, 앞 2글자 다르면 false",
			given1:      "abc", given2: "bbc",
			want: false,
		},
	} {
		t.Run(fmt.Sprintf("%s %s, %s", tt.description, tt.given1, tt.given2), func(t *testing.T) {
			got := equalPrefix(tt.given1, tt.given2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPalindromeIndexPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		file, err := os.Open("./test_data/palindrome_index.csv")
		assert.NoError(t, err)
		defer func() {
			err := file.Close()
			assert.NoError(t, err)
		}()

		r := csv.NewReader(bufio.NewReader(file))
		rows, err := r.ReadAll()
		assert.NoError(t, err)

		var arr []string
		for _, row := range rows {
			arr = append(arr, row[0])
		}

		want := []int32{
			16722,
			76661,
			10035,
			57089,
			46674,
		}

		assert.Len(t, arr, 5)
		for i, s := range arr {
			got := palindromeIndex(s)
			assert.EqualValues(t, want[i], got)
		}
		return true
	}, time.Second, time.Millisecond*100, "시간 초과")
}
