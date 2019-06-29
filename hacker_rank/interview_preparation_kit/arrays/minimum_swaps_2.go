package arrays

func find(arr []int32, target int32, offset int, l int) (int, bool) {
	for i := offset; i < l; i++ {
		if arr[i] == target {
			return i, true
		}
	}

	return -1, false
}

func minimumSwaps(arr []int32) int32 {
	var cnt int32
	length := len(arr)
	for i, n := range arr {
		idx := i + 1
		if int32(idx) != n {
			pos, exist := find(arr, int32(idx), i, length)
			if !exist {
				return -1
			}

			arr[i], arr[pos] = arr[pos], arr[i]
			cnt++
		}
	}

	return cnt
}
