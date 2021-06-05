package str

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	t.Parallel()

	t.Log("Sort 함수는 하나의 문자열을 잘 정렬한다.")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		given string
		want  string
	}{
		{"dcba", "abcd"},
		{"ffbbaa", "aabbff"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Sort(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSortAdapter(t *testing.T) {
	t.Parallel()

	t.Log("SortAdapter 구조체는 문자열 슬라이스를 잘 정렬한다.")

	for i, tt := range []struct {
		given []string
		want  []string
	}{
		{
			[]string{"a", "d", "b", "c"},
			[]string{"a", "b", "c", "d"},
		},
		{
			[]string{"a", "d", "bc", "ba"},
			[]string{"a", "ba", "bc", "d"},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := make([]string, len(tt.given))
			copy(got, tt.given)
			sort.Sort(SortAdapter(got))
			assert.Equal(t, tt.want, got)
		})
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
