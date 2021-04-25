package strings

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeIndex(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/palindrome-index/problem")

	for _, tc := range []struct {
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
		t.Run(tc.s, func(t *testing.T) {
			got := palindromeIndex(tc.s)
			assert.EqualValues(t, tc.want, got)
		})
	}
}

func TestEqualPrefix(t *testing.T) {
	t.Run("1글자 문자열 두 개", func(t *testing.T) {
		t.Run("같으면 true 반환", func(t *testing.T) {
			got := equalPrefix("a", "a")
			assert.True(t, got)
		})

		t.Run("다르면 false 반환", func(t *testing.T) {
			got := equalPrefix("a", "b")
			assert.False(t, got)
		})
	})

	t.Run("2글자 이상의 길이의 문자열 두 개의 앞 2글자", func(t *testing.T) {
		t.Run("같으면 true를 반환.", func(t *testing.T) {

			got := equalPrefix("abcd", "abab")
			assert.True(t, got)
		})

		t.Run("다르면 false를 반환한다.", func(t *testing.T) {
			got := equalPrefix("abc", "bbc")
			assert.False(t, got)
		})
	})
}

func TestPalindromeIndexPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		file, err := os.Open("./test_data/palindrome_index.csv")
		assert.NoError(t, err)
		defer func() {
			_ = file.Close()
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
