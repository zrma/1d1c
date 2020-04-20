package strings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/bear-and-steady-gene/problem", func() {
	It("문제를 풀었다", func() {
		Expect(steadyGene("GAAATAAA")).Should(BeNumerically("==", 5))
		Expect(steadyGene("TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC")).Should(BeNumerically("==", 5))
	})
})
