package lv_0_easy

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cnt := 1
	prev := nums[0]
	for _, n := range nums[1:] {
		if n != prev {
			nums[cnt] = n
			cnt++
			prev = n
		}
	}
	return cnt
}
