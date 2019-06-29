package warmup

func jumpingOnClouds(c []int32) int32 {
	var total int32
	var before int32
	length := len(c)
	for cur, v := range c {
		if v == 1 {
			step := jump(before, int32(cur-1))
			if step < 0 {
				return -1
			}

			total += step + 1
			before = int32(cur) + 1
		} else if cur == (length - 1) {
			total += (int32(cur) - before + 1) / 2
		}
	}

	return total
}

func jump(before, after int32) int32 {
	diff := after - before
	if diff < 0 {
		return -1
	}

	return (diff + 1) / 2
}
