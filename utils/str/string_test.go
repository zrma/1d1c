package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	t.Parallel()

	t.Log("Sort 함수는 하나의 문자열을 잘 정렬한다.")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		given string
		want  string
	}{
		{"dcba", "abcd"},
		{"ffbbaa", "aabbff"},
	} {
		got := Sort(tt.given)
		assert.Equal(t, tt.want, got)
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()

	t.Run("문자열을 잘 뒤집는다", func(t *testing.T) {
		const (
			given = "abc"
			want  = "cba"
		)
		got := Reverse(given)
		assert.Equal(t, want, got)
	})

	t.Run("빈 문자열에 오류가 발생하지 않는다.", func(t *testing.T) {
		assert.NotPanics(t, func() {
			const emptyStr = ""
			got := Reverse(emptyStr)
			assert.Equal(t, emptyStr, got)
		})
	})
}
