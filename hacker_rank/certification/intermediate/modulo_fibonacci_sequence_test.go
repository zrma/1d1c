package intermediate

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestModuloFibonacciSequence(t *testing.T) {
	for _, tt := range []struct {
		skip, total int32
		want        []int
	}{
		{
			0, 6,
			[]int{1, 2, 3, 5, 8, 13},
		},
		{
			50, 4,
			[]int{
				951280099,
				316291173,
				267571272,
				583862445,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.skip, tt.total), func(t *testing.T) {
			resultChan := make(chan int)
			requestChan := make(chan bool)
			go ModuloFibonacciSequence(requestChan, resultChan)

			var got []int
			for i := int32(0); i < tt.skip+tt.total; i++ {
				start := time.Now().UnixNano()
				requestChan <- true
				res := <-resultChan
				if i < tt.skip {
					continue
				}
				end := time.Now().UnixNano()
				timeDiff := (end - start) / 1000000
				if timeDiff < 3 {
					fmt.Println("Rate is too high")
					break
				}
				got = append(got, res)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
