package tutorial30daysofcode

import "fmt"

func dataType(i1, i2 uint64, d1, d2 float64, s1, s2 string) {
	fmt.Println(i1 + i2)
	fmt.Printf("%.1f\n", d1+d2)
	fmt.Println(s1 + s2)
}
