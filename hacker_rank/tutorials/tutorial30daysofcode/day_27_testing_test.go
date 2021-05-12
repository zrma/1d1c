package tutorial30daysofcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumIndex(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-testing/problem")

	t.Run("emptyArray() 함수는 비어 있는 슬라이스를 잘 반환한다.", func(t *testing.T) {
		const want = invalidIdx
		given := emptyArray()

		got, err := minimumIndex(given)
		assert.Error(t, err)
		assert.EqualError(t, err, errMsg)
		assert.Equal(t, want, got)
	})

	t.Run("uniqueValues() 함수는 유니크한 두 개 이상의 값을 담고 있는 슬라이스를 반환한다.", func(t *testing.T) {
		seq := uniqueValues()
		assert.GreaterOrEqual(t, len(seq), 2)

		m := make(map[int]interface{})
		for _, num := range seq {
			_, ok := m[num]
			assert.False(t, ok)
			m[num] = nil
		}

		got, err := minimumIndex(seq)
		assert.NoError(t, err)
		assert.Equal(t, 0, got, "해당 풀이에서는 minimumIndex = 0")
	})

	t.Run("exactlyTwoDifferentMinimums() 함수는 가장 작은 정수 두 개와 임의의 다른 더 큰 0개 이상의 정수로 이뤄진 슬라이스를 반환한다.", func(t *testing.T) {
		seq := exactlyTwoDifferentMinimums()
		assert.GreaterOrEqual(t, len(seq), 2)

		tmp := make([]int, len(seq))
		copy(tmp, seq)
		sort.Slice(tmp, func(i, j int) bool {
			return i > j
		})
		assert.Equal(t, tmp[0], tmp[1])
		assert.True(t, len(tmp) == 2 || tmp[1] < tmp[2])

		got, err := minimumIndex(seq)
		assert.NoError(t, err)
		assert.Equal(t, 1, got, "해당 풀이에서는 minimumIndex = 1")
	})
}
