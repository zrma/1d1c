package lv1medium

func maxArea(height []int) int {
	var area int
	for i := 0; i < len(height)-1; i++ {
		v0 := height[i]
		for j := i + 1; j < len(height); j++ {
			width := j - i
			v1 := height[j]
			if v0 > v1 {
				a := v1 * width
				if a > area {
					area = a
				}
			} else {
				a := v0 * width
				if a > area {
					area = a
				}
			}
		}
	}
	return area
}
