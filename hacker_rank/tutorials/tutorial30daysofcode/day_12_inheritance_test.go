package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-inheritance/problem", func() {
	//noinspection SpellCheckingInspection
	It("문제를 풀었다", func() {
		s := newStudent("Heraldo", "Memelli", 8135627, []int{100, 80})
		Expect(s.printPerson()).Should(Equal("Name: Memelli, Heraldo\nID: 8135627"))
		Expect(s.calculate()).Should(Equal("O"))

		s.testScores = []int{80}
		Expect(s.calculate()).Should(Equal("E"))
		s.testScores = []int{70}
		Expect(s.calculate()).Should(Equal("A"))
		s.testScores = []int{55}
		Expect(s.calculate()).Should(Equal("P"))
		s.testScores = []int{40}
		Expect(s.calculate()).Should(Equal("D"))
		s.testScores = []int{30}
		Expect(s.calculate()).Should(Equal("T"))
	})
})
