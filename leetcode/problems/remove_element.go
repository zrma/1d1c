package problems

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	cnt := 0
	for _, n := range nums {
		if n != val {
			nums[cnt] = n
			cnt++
		}
	}
	return cnt
}
