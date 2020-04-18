package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMorganAndString(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/morgan-and-a-string/problem", func(t *testing.T) {
		for _, data := range []struct {
			jack, daniel string
			expected     string
		}{
			{"ACA", "BCF", "ABCACF"},
			{"JACK", "DANIEL", "DAJACKNIEL"},
			{"ABACABA", "ABACABA", "AABABACABACABA"},
			{"AAAAB", "AAAAC", "AAAAAAAABC"},
			{"AAAAC", "AAAAB", "AAAAAAAABC"},
		} {
			actual := morganAndString(data.jack, data.daniel)
			assert.Equal(t, data.expected, actual)
		}
	})
}
