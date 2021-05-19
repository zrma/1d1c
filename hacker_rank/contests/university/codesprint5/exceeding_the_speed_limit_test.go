package codesprint5

import (
	"fmt"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestExceedingTheSpeedLimit(t *testing.T) {
	for _, tt := range []struct {
		given int32
		want  string
	}{
		{100, "3000 Warning"},
		{140, "25000 License removed"},
		{85, "0 No punishment"},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			want := []string{tt.want}
			err := utils.PrintTest(func() {
				exceedingTheSpeedLimit(tt.given)
			}, want)
			assert.NoError(t, err)
		})
	}
}
