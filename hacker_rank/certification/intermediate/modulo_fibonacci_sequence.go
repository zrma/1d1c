package intermediate

import (
	"time"
)

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	var x0 = 1
	var x1 = 1
	const mod = 1000000000
	for req := range requestChan {
		time.Sleep(time.Millisecond * 10)
		if !req {
			return
		}
		resultChan <- x1
		tmp := x1
		x1 += x0
		x1 %= mod
		x0 = tmp
	}
}
