package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteadyGene(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/bear-and-steady-gene/problem", func(t *testing.T) {
		assert.Equal(t, int32(5), steadyGene("GAAATAAA"))
		assert.Equal(t, int32(5), steadyGene("TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC"))
	})
}
