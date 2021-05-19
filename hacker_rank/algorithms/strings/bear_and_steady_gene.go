package strings

import (
	"math"

	"1d1go/utils/integer"
)

// 넘치는 문자열을 넘치는 만큼만 포함하는 최소 구간을 찾는다.
func steadyGene(gene string) int32 {
	length := len(gene)
	increase, decrease, validate := calculators(length)
	for _, g := range gene {
		increase(byte(g))
	}

	var begin, end int32
	var min int32 = math.MaxInt32
	for end < int32(length-1) {
		decrease(gene[end])
		end++

		for validate() {
			min = integer.MinInt32(min, end-begin)
			increase(gene[begin])
			begin++
		}
	}

	return min
}

type validateFunc func() bool
type increaseFunc func(byte)
type decreaseFunc func(byte)

func calculators(l int) (increaseFunc, decreaseFunc, validateFunc) {
	cnt := l / 4
	var arr = [4]int{-cnt, -cnt, -cnt, -cnt}

	acc := func(b byte, step int) {
		switch b {
		case 'A':
			arr[0] += step
		case 'C':
			arr[1] += step
		case 'T':
			arr[2] += step
		case 'G':
			arr[3] += step
		}
	}

	return func(b byte) {
			acc(b, 1)
		}, func(b byte) {
			acc(b, -1)
		}, func() bool {
			for _, n := range arr {
				if n > 0 {
					return false
				}
			}
			return true
		}
}
