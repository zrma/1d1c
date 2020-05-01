package lv0easy

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}
	return result
}

func twoSum2InputArrayIsSorted(nums []int, target int) []int {
	result := twoSum(nums, target)
	for i := range result {
		result[i]++
	}
	return result
}
