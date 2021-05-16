package lv0easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Log("https://leetcode.com/problems/two-sum/")

	const target = 9
	var (
		nums = []int{2, 7, 11, 15}
		want = []int{0, 1}
	)
	got := twoSum(nums, target)
	assert.Equal(t, want, got)
}

func TestTwoSum2InputArrayIsSorted(t *testing.T) {
	t.Log("https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/")

	var (
		nums   = []int{2, 7, 11, 15}
		target = 9
		want   = []int{1, 2}
	)
	got := twoSum2InputArrayIsSorted(nums, target)
	assert.Equal(t, want, got)
}
