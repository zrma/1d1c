package implementation

import (
	"sort"
)

func getTotalX(arr1, arr2 []int32) int32 {
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	begin := arr1[len(arr1)-1]
	end := arr2[0]

	cond1 := func(num int32, arr []int32) bool {
		for _, n := range arr {
			if num%n != 0 {
				return false
			}
		}
		return true
	}
	cond2 := func(num int32, arr []int32) bool {
		for _, n := range arr {
			if n%num != 0 {
				return false
			}
		}
		return true
	}

	var cnt int32
	for n := begin; n <= end; n++ {
		if cond1(n, arr1) && cond2(n, arr2) {
			cnt++
		}
	}
	return cnt
}
