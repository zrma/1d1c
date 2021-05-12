package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMorganAndString(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/morgan-and-a-string/problem")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		jack, daniel string
		want         string
	}{
		{"ACA", "BCF", "ABCACF"},
		{"JACK", "DANIEL", "DAJACKNIEL"},
		{"ABACABA", "ABACABA", "AABABACABACABA"},
		{"AAAAB", "AAAAC", "AAAAAAAABC"},
		{"AAAAC", "AAAAB", "AAAAAAAABC"},
	} {
		t.Run(fmt.Sprintf("%d - %s %s", i, tt.jack, tt.daniel), func(t *testing.T) {
			got := morganAndString(tt.jack, tt.daniel)
			assert.Equal(t, tt.want, got)
		})
	}
}
