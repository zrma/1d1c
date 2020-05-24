package lv2hard

import "math"

func findTwoValuesNearMid(nums1 []int, nums2 []int, totLen int) (prev, curr int) {
	target := totLen / 2
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
	totLen := len(nums1) + len(nums2)
	val1, val2 := findTwoValuesNearMid(nums1, nums2, totLen)
	if totLen%2 == 0 {
		return float64(val1+val2) / 2
	}
	return float64(val2)
}

func mergeSortedArray(nums1 []int, nums2 []int) []int {
	totLen := len(nums1) + len(nums2)
	totNums := make([]int, totLen)

	var i, j, cur int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			totNums[cur] = nums1[i]
			cur++
			i++
		} else {
			totNums[cur] = nums2[j]
			cur++
			j++

		}
	}
	for i < len(nums1) {
		totNums[cur] = nums1[i]
		cur++
		i++

	}
	for j < len(nums2) {
		totNums[cur] = nums2[j]
		cur++
		j++

	}
	return totNums
}

func findMedianSortedArraysWithMerge(nums1 []int, nums2 []int) float64 {
	totNums := mergeSortedArray(nums1, nums2)
	if len(totNums)%2 == 0 {
		return float64(totNums[len(totNums)/2-1]+totNums[len(totNums)/2]) / 2
	} else {
		return float64(totNums[len(totNums)/2])
	}

}

type side struct {
	nums1 int
	nums2 int
}

func findBoundary(nums1 []int, nums2 []int, totCnt int) (leftMax, rightMin side) {
	// range of binary search target in nums1
	low, high := 0, len(nums1)

	for low <= high {
		// nums1:  1 2 4 | 5 6 7
		// nums2:  3 5   | 8 9
		//
		// totCnt := 6 + 4
		// cut1 := (0 + 6) / 2 = 3
		// cut2 := (totCnt+1) / 2 - cut1 = 5 - 3 = 2
		//
		// 총 10개이므로 좌우측 중 한 쪽에서 5개씩 가져갈 수 있음
		// totCnt 에 더해주는 +1 은 인덱스 계산을 편하게 하기 위한 sugar
		// 좌측의 nums1에서 먼저 3개를 가져가면 nums2는 2개만 가져갈 수 있음
		//
		// it makes both side(left and right) being equal count
		cut1 := (low + high) / 2
		cut2 := (totCnt+1)/2 - cut1

		if cut1 < 1 {
			leftMax.nums1 = math.MinInt32
		} else {
			leftMax.nums1 = nums1[cut1-1]
		}

		if cut1 >= len(nums1) {
			rightMin.nums1 = math.MaxInt32
		} else {
			rightMin.nums1 = nums1[cut1]
		}

		if cut2 < 1 {
			leftMax.nums2 = math.MinInt32
		} else {
			leftMax.nums2 = nums2[cut2-1]
		}
		if cut2 >= len(nums2) {
			rightMin.nums2 = math.MaxInt32
		} else {
			rightMin.nums2 = nums2[cut2]
		}

		if leftMax.nums1 <= rightMin.nums2 && leftMax.nums2 <= rightMin.nums1 {
			break
		} else {
			if leftMax.nums1 >= rightMin.nums2 {
				high = cut1 - 1
			} else {
				low = cut1 + 1
			}
		}
	}
	return
}

func findMedianSortedArraysWithBinSearch(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	totCnt := len(nums1) + len(nums2)
	leftMax, rightMin := findBoundary(nums1, nums2, totCnt)

	leftMaxVal := math.Max(float64(leftMax.nums1), float64(leftMax.nums2))
	if totCnt%2 == 0 {
		rightMinVal := math.Min(float64(rightMin.nums1), float64(rightMin.nums2))
		return (leftMaxVal + rightMinVal) / 2
	} else {
		return leftMaxVal
	}
}

func findMedianSortedArraysWithBinSearchSolution(nums1 []int, nums2 []int) float64 {
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
