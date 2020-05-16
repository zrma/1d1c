package lv2hard

func findTwoValuesNearMid(nums1 []int, nums2 []int, totalLength int) (prev, curr int) {
	target := totalLength / 2
	var idx1, idx2, idxTot int
	for idx1 < len(nums1) && idx2 < len(nums2) {
		prev = curr
		if nums1[idx1] < nums2[idx2] {
			curr = nums1[idx1]
			idxTot++
			idx1++
		} else {
			curr = nums2[idx2]
			idxTot++
			idx2++
		}
		if idxTot == target+1 {
			return
		}
	}
	for idx1 < len(nums1) {
		prev = curr
		curr = nums1[idx1]
		idxTot++
		idx1++
		if idxTot == target+1 {
			return
		}
	}
	for idx2 < len(nums2) {
		prev = curr
		curr = nums2[idx2]
		idxTot++
		idx2++
		if idxTot == target+1 {
			return
		}
	}
	return
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	val1, val2 := findTwoValuesNearMid(nums1, nums2, totalLength)
	if totalLength%2 == 0 {
		return float64(val1+val2) / 2
	}
	return float64(val2)
}
