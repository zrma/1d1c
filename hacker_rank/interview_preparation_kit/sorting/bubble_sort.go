package sorting

import (
	"fmt"
)

func countSwaps(a []int32) {
	swapCnt := 0
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapCnt++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", swapCnt)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[n-1])
}
