package arrays

func rotLeft(a []int32, d int32) []int32 {
	length := int32(len(a))

	if (d < 1) || (d%length == 0) {
		return a
	}

	if d > length {
		d = d % length
	}

	result := a[d:]
	return append(result, a[:d]...)
}
