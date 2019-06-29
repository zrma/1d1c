package arrays

func rotLeft(a []int32, d int32) []int32 {
	l := int32(len(a))

	if (d < 1) || (d%l == 0) {
		return a
	}

	if d > l {
		d = d % l
	}

	result := a[d:]
	return append(result, a[:d]...)
}
