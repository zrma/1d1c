package lv2hard

import "math"

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

func findMedianSortedArraysWithMerge(nums1 []int, nums2 []int) float64 {
	total := mergeSortedArray(nums1, nums2)
	if len(total)%2 == 0 {
		return float64(total[len(total)/2-1]+total[len(total)/2]) / 2
	} else {
		return float64(total[len(total)/2])
	}

}

func findMedianSortedArraysWithBinSearch(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	half := (len(nums1) + len(nums2) - 1) / 2

	l, r := 0, len(nums1)
	i := (l + r) / 2
	j := half - i
	for l < r {
		if nums1[i] < nums2[j] {
			l = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j+1] {
			r = i
		} else {
			break
		}
		i = (l + r) / 2
		j = half - i
	}

	var medianLeft, medianRight float64
	if j < 0 {
		medianLeft = float64(nums1[i-1])
	} else if i < 1 {
		medianLeft = float64(nums2[j])
	} else {
		medianLeft = math.Max(float64(nums1[i-1]), float64(nums2[j]))
	}

	if (len(nums1)+len(nums2))%2 == 1 {
		return medianLeft
	} else if i == len(nums1) {
		medianRight = float64(nums2[j+1])
	} else if j+1 == len(nums2) {
		medianRight = float64(nums1[i])
	} else {
		medianRight = math.Min(float64(nums1[i]), float64(nums2[j+1]))
	}
	return (medianLeft + medianRight) / 2
}
