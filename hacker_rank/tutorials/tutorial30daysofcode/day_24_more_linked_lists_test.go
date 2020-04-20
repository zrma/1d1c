package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-linked-list-deletion/problem", func() {
	It("문제를 풀었다", func() {
		var list linkedList
		for _, num := range []int{1, 2, 2, 3, 3, 4} {
			list.head = list.insert(num)
		}

		list.head = removeDuplicates(list.head)
		err := utils.PrintTest(func() {
			list.display()
		}, []string{
			"1", "2", "3", "4",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("조금 더 까다로운 테스트 케이스", func() {
		var list linkedList
		for _, num := range []int{1, 2, 2, 2, 3, 3, 3, 3, 4, 4} {
			list.head = list.insert(num)
		}

		list.head = removeDuplicates(list.head)
		err := utils.PrintTest(func() {
			list.display()
		}, []string{
			"1", "2", "3", "4",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
