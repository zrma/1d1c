package tutorial30daysofcode

import "fmt"

func loop(n int32) {
	var i int32
	for i = 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
