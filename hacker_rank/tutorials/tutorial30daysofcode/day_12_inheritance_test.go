package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStudent(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-inheritance/problem")

	for _, tt := range []struct {
		given int
		want  string
	}{
		{80, "E"},
		{70, "A"},
		{55, "P"},
		{40, "D"},
		{30, "T"},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			//noinspection SpellCheckingInspection
			const (
				firstName      = "Heraldo"
				lastName       = "Memelli"
				identification = 8135627
			)
			s := newStudent(firstName, lastName, identification, []int{100, 80})

			gotPerson := s.printPerson()
			wantPerson := fmt.Sprintf("Name: %s, %s\nID: %d", lastName, firstName, identification)
			assert.Equal(t, wantPerson, gotPerson)

			gotScore := s.calculate()
			const wantScore = "O"
			assert.Equal(t, wantScore, gotScore)

			s.testScores = []int{tt.given}
			got := s.calculate()
			assert.Equal(t, tt.want, got)
		})
	}
}
