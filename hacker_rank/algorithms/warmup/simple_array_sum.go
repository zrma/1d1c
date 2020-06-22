package warmup

func simpleArraySum(arr []int32) int32 {
	var sum int32
	for _, n := range arr {
		sum += n
	}

	return sum
}
