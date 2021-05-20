package basic

func countBits(num uint32) int32 {
	var res int32
	for num > 0 {
		res += int32(num % 2)
		num /= 2
	}
	return res
}
