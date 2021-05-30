package intermediate

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			defer close(resultChan)
			requestChan := make(chan bool)
			defer close(requestChan)

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
				require.Greater(t, timeDiff, int64(3), "Rate is too high")

				got = append(got, res)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModuloFibonacciSequenceReqFalse(t *testing.T) {
	resultChan := make(chan int)
	requestChan := make(chan bool)
	go ModuloFibonacciSequence(requestChan, resultChan)

	waitChan := make(chan interface{})
	go func() {
		defer close(waitChan)

		start := time.Now().UnixNano()
		got := <-resultChan
		require.Zero(t, got)
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		assert.Greater(t, timeDiff, int64(3))
	}()
	requestChan <- false

	time.Sleep(time.Millisecond * 10)

	close(requestChan)
	close(resultChan)

	<-waitChan
}
