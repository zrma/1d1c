package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/anagram/problem", func(t *testing.T) {
		for _, data := range []struct {
			input    string
			expected int32
		}{
			{"", -1},
			{"aaabbb", 3},
			{"ab", 1},
			{"abc", -1},
			{"mnop", 2},
			{"xyyx", 0},
			{"xaxbbbxx", 1},
			{"hhpddlnnsjfoyxpciioigvjqzfbpllssuj", 10},
			{"xulkowreuowzxgnhmiqekxhzistdocbnyozmnqthhpievvlj", 13},
			{"dnqaurlplofnrtmh", 5},
			{"aujteqimwfkjoqodgqaxbrkrwykpmuimqtgulojjwtukjiqrasqejbvfbixnchzsahpnyayutsgecwvcqngzoehrmeeqlgknnb", 26},
			{"lbafwuoawkxydlfcbjjtxpzpchzrvbtievqbpedlqbktorypcjkzzkodrpvosqzxmpad", 15},
			{"drngbjuuhmwqwxrinxccsqxkpwygwcdbtriwaesjsobrntzaqbe", -1},
			{"ubulzt", 3},
			{"vxxzsqjqsnibgydzlyynqcrayvwjurfsqfrivayopgrxewwruvemzy", 13},
			{"xtnipeqhxvafqaggqoanvwkmthtfirwhmjrbphlmeluvoa", 13},
			{"gqdvlchavotcykafyjzbbgmnlajiqlnwctrnvznspiwquxxsiwuldizqkkaawpyyisnftdzklwagv", -1},
		} {
			assert.Equal(t, data.expected, anagram(data.input))
		}
	})
}
