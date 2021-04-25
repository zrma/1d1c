package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteadyGene(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/bear-and-steady-gene/problem")

	got := steadyGene("GAAATAAA")
	assert.EqualValues(t, 5, got)

	got = steadyGene("TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC")
	assert.EqualValues(t, 5, got)
}
