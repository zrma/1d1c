package strings

import (
	"github.com/zrma/going/utils/integer"
	"math"
)

type validateFunc func() bool
type increaseFunc func(byte)
type decreaseFunc func(byte)

func calculators(l int) (increaseFunc, decreaseFunc, validateFunc) {
	cnt := l / 4
	var arr = [4]int{-cnt, -cnt, -cnt, -cnt}

	calcArr := func(b byte, margin int) {
		switch b {
		case 'A':
			arr[0] += margin
		case 'C':
			arr[1] += margin
		case 'T':
			arr[2] += margin
		case 'G':
			arr[3] += margin
		}
	}

	return func(b byte) {
			calcArr(b, 1)
		}, func(b byte) {
			calcArr(b, -1)
		}, func() bool {
			for _, num := range arr {
				if num > 0 {
					return false
				}
			}
			return true
		}
}

// 넘치는 문자열을 넘치는 만큼만 포함하는 최소 구간을 찾는다.
func steadyGene(gene string) int32 {
	length := len(gene)
	increase, decrease, validate := calculators(length)
	for _, g := range gene {
		increase(byte(g))
	}

	var begin, end int32
	var minRange int32 = math.MaxInt32
	for end < int32(length-1) {
		decrease(gene[end])
		end++

		for validate() {
			minRange = integer.MinInt32(minRange, end-begin)
			increase(gene[begin])
			begin++
		}
	}

	return minRange
}
