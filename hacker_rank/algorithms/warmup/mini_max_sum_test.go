package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestMiniMaxSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/mini-max-sum/problem")

	want := []string{"10 14"}
	got, err := utils.GetPrinted(func() {
		miniMaxSum([]int64{1, 2, 3, 4, 5})
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
