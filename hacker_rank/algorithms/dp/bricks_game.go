package dp

import (
	"github.com/zrma/going/utils/integer"
)

// 메모이제이션을 위해 캐시를 사용한다.
// 그러지 않으면 성능 테스트 케이스 타임아웃에 걸린다.
type brickCache map[int]innerCache

type innerCache map[int]int64

func (cache *brickCache) get(begin, end int) (int64, bool) {
	if inner, ok := (*cache)[begin]; ok {
		result, ok := inner[end]
		return result, ok
	}

	(*cache)[begin] = make(innerCache)
	return 0, false
}

func playSubGame(arr []int32, begin, end int, cache brickCache) int64 {
	if result, ok := cache.get(begin, end); ok {
		return result
	}

	// 재귀 함수 탈출 조건
	if end-begin <= 3 {
		var sum int64
		for i := 0; begin+i <= end; i++ {
			if i == 3 {
				break
			}
			sum += int64(arr[begin+i])
		}

		cache[begin][end] = sum
		return sum
	}

	// 나는 최대한 점수를 많이 내려고 한다.
	result := integer.MaxInt64(
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		integer.MinInt64(
			int64(arr[begin])+playSubGame(arr, begin+2, end, cache),
			int64(arr[begin])+playSubGame(arr, begin+3, end, cache),
			int64(arr[begin])+playSubGame(arr, begin+4, end, cache),
		),
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		integer.MinInt64(
			int64(arr[begin])+int64(arr[begin+1])+playSubGame(arr, begin+3, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+playSubGame(arr, begin+4, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+playSubGame(arr, begin+5, end, cache),
		),
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		integer.MinInt64(
			int64(arr[begin])+int64(arr[begin+1])+int64(arr[begin+2])+playSubGame(arr, begin+4, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+int64(arr[begin+2])+playSubGame(arr, begin+5, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+int64(arr[begin+2])+playSubGame(arr, begin+6, end, cache),
		),
	)
	cache[begin][end] = result
	return result
}

func bricksGame(arr []int32) int64 {
	cache := make(brickCache)
	return playSubGame(arr, 0, len(arr)-1, cache)
}
