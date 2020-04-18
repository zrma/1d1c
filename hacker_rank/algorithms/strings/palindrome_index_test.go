package strings

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/zrma/going/utils"
)

func TestPalindromeIndex(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/palindrome-index/problem", func(t *testing.T) {
		for _, data := range []struct {
			s        string
			expected int32
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
			actual := palindromeIndex(data.s)
			assert.Equal(t, data.expected, actual)
		}
	})

	t.Run("equalPrefix 함수는", func(t *testing.T) {
		t.Run("1글자 문자열 두 개가", func(t *testing.T) {
			{
				actual := equalPrefix("a", "a")
				assert.True(t, actual, "같으면 true 반환")
			}
			{
				actual := equalPrefix("a", "b")
				assert.False(t, actual, "다르면 false 반환")
			}
		})

		//noinspection SpellCheckingInspection
		t.Run("2글자 이상의 길이의 문자열 두 개의 앞 2글자", func(t *testing.T) {
			{
				actual := equalPrefix("abcd", "abab")
				assert.True(t, actual, "같으면 true 반환")
			}
			{
				actual := equalPrefix("abc", "bbc")
				assert.False(t, actual, "다르면 false 반환")
			}
		})
	})

	t.Run("performance measure", func(t *testing.T) {
		RunUntil(t, func(done Done) {
			defer close(done)

			file, err := os.Open("./test_data/palindrome_index.csv")
			assert.NoError(t, err)
			defer file.Close()

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			assert.NoError(t, err)

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

			assert.Equal(t, 5, len(arr))
			for i, s := range arr {
				actual := palindromeIndex(s)
				assert.Equal(t, expected[i], actual)
			}
		}, 3)
	})
}
