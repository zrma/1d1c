package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	for _, tt := range []struct {
		given uint32
		want  int32
	}{
		{127, 7},
		{10, 2},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			got := countBits(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
