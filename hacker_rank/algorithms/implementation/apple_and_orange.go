package implementation

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	containFrom := func(origin int32) func(int32) bool {
		return func(distance int32) bool {
			return (s <= origin+distance) && (origin+distance <= t)
		}
	}

	fmt.Println(count(apples, containFrom(a)))
	fmt.Println(count(oranges, containFrom(b)))
}

func count(distances []int32, contain func(int32) bool) (sum int32) {
	for _, distance := range distances {
		if contain(distance) {
			sum++
		}
	}
	return
}
