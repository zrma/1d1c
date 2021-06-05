package tutorial30daysofcode

import "fmt"

func printReverse(arr []int32) {
	for i := len(arr); i > 0; i-- {
		fmt.Printf("%d", arr[i-1])
		if i > 1 {
			fmt.Printf(" ")
		}
	}
}
