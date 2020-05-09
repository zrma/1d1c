package sorting

import (
	"sort"
)

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	var cnt int32
	for _, price := range prices {
		k -= price
		if k < 0 {
			break
		}
		cnt++
	}
	return cnt
}
