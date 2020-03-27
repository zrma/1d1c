package tutorial30daysofcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumIndex(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-testing/problem", func(t *testing.T) {
		t.Run("emptyArray() 함수는 비어 있는 슬라이스를 잘 반환한다.", func(t *testing.T) {
			actual, err := minimumIndex(emptyArray())
			assert.Error(t, err)

			const expected = -1
			assert.Equal(t, expected, actual)
		})

		t.Run("uniqueValues() 함수는 유니크한 값만 담고 있는 슬라이스를 반환한다.", func(t *testing.T) {
			seq := uniqueValues()
			assert.Len(t, seq, 2)

			m := make(map[int]interface{})
			for _, num := range seq {
				_, ok := m[num]
				assert.False(t, ok)
				m[num] = nil
			}
			actual, err := minimumIndex(seq)
			assert.NoError(t, err)
			assert.Empty(t, actual)
		})

		t.Run("exactlyTwoDifferentMinimums() 함수는 가장 작은 정수 두 개와 임의의 다른 더 큰 정수들로 이뤄진 슬라이스를 반환한다.", func(t *testing.T) {
			seq := exactlyTwoDifferentMinimums()
			assert.Len(t, seq, 3)

			tmp := make([]int, len(seq))
			copy(tmp, seq)
			sort.Slice(tmp, func(i, j int) bool {
				return i > j
			})

			assert.Equal(t, tmp[0], tmp[1])
			assert.True(t, len(tmp) == 2 || tmp[1] < tmp[2])

			actual, err := minimumIndex(seq)
			assert.NoError(t, err)

			const expected = 1
			assert.Equal(t, expected, actual)
		})
	})
}
