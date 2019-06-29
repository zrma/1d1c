package arrays

func rotLeft(arr []int32, n int32) []int32 {
	length := int32(len(arr))

	if (n < 1) || (n%length == 0) {
		return arr
	}

	if n > length {
		n = n % length
	}

	result := arr[n:]
	return append(result, arr[:n]...)
}
