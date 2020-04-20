package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-queues-stacks/problem", func() {
	Context("문제를 풀었다", func() {
		It("queue 구조체는 큐 자료구조의 기능을 충실히 수행한다.", func() {
			q := queue{}

			q.enqueue("a")
			q.enqueue("b")
			q.enqueue("c")

			actual, ok := q.dequeue()
			Expect(ok).Should(BeTrue())
			Expect(actual).Should(Equal("a"))

			actual, ok = q.dequeue()
			Expect(ok).Should(BeTrue())
			Expect(actual).Should(Equal("b"))

			actual, ok = q.dequeue()
			Expect(ok).Should(BeTrue())
			Expect(actual).Should(Equal("c"))

			_, ok = q.dequeue()
			Expect(ok).Should(BeFalse())
		})

		It("stack 구조체는 스택 자료구조의 기능을 충실히 수행한다.", func() {
			s := stack{}

			s.push("a")
			s.push("b")
			s.push("c")

			actual, ok := s.pop()
			Expect(ok).Should(BeTrue())
			Expect(actual).Should(Equal("c"))

			actual, ok = s.pop()
			Expect(ok).Should(BeTrue())
			Expect(actual).Should(Equal("b"))

			actual, ok = s.pop()
			Expect(ok).Should(BeTrue())
			Expect(actual).Should(Equal("a"))

			_, ok = s.pop()
			Expect(ok).Should(BeFalse())
		})
	})
})
