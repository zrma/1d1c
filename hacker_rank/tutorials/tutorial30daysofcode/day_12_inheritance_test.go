package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStudent(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/30-inheritance/problem", func(t *testing.T) {
		s := newStudent("Heraldo", "Memelli", 8135627, []int{100, 80})
		assert.Equal(t, s.printPerson(), "Name: Memelli, Heraldo\nID: 8135627")
		assert.Equal(t, s.calculate(), "O")

		for score, expected := range map[int]string{
			80: "E", 70: "A", 55: "P", 40: "D", 30: "T",
		} {
			s.testScores = []int{score}
			actual := s.calculate()
			assert.Equal(t, expected, actual)
		}
	})
}
