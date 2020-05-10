package lv2hard

func mergeSortedArray(nums1 []int, nums2 []int) []int {
	totalLength := len(nums1) + len(nums2)
	total := make([]int, totalLength)

	var i, j, cur int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			total[cur] = nums1[i]
			cur++
			i++
		} else {
			total[cur] = nums2[j]
			cur++
			j++
		}
	}
	for i < len(nums1) {
		total[cur] = nums1[i]
		cur++
		i++
	}
	for j < len(nums2) {
		total[cur] = nums2[j]
		cur++
		j++
	}
	return total
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := mergeSortedArray(nums1, nums2)
	if len(total)%2 == 0 {
		return float64(total[len(total)/2-1]+total[len(total)/2]) / 2
	} else {
		return float64(total[len(total)/2])
	}
}
