package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestMoreException(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-more-exceptions/problem")

	for _, tt := range []struct {
		n, p int64
		want string
	}{} {
		t.Run(fmt.Sprintf("%d %d", tt.n, tt.p), func(t *testing.T) {
			err := utils.PrintTest(func() {
				moreException(tt.n, tt.p)
			}, []string{
				tt.want,
			})
			assert.NoError(t, err)
		})
	}
}
