package tutorial30daysofcode

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-testing/problem", func() {
	Context("문제를 풀었다", func() {
		It("emptyArray() 함수는 비어 있는 슬라이스를 잘 반환한다.", func() {
			actual, err := minimumIndex(emptyArray())
			Expect(err).Should(HaveOccurred())
			Expect(actual).Should(BeNumerically("==", -1))
		})

		It("uniqueValues() 함수는 유니크한 값만 담고 있는 슬라이스를 반환한다.", func() {
			seq := uniqueValues()
			Expect(len(seq)).Should(BeNumerically(">=", 2))

			m := make(map[int]interface{})
			for _, num := range seq {
				_, ok := m[num]
				Expect(ok).Should(BeFalse())
				m[num] = nil
			}
			actual, err := minimumIndex(seq)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(actual).Should(BeNumerically("==", 0))
		})

		It("exactlyTwoDifferentMinimums() 함수는 가장 작은 정수 두 개와 임의의 다른 더 큰 정수들로 이뤄진 슬라이스를 반환한다.", func() {
			seq := exactlyTwoDifferentMinimums()
			Expect(len(seq)).Should(BeNumerically(">=", 2))

			tmp := make([]int, len(seq))
			copy(tmp, seq)
			sort.Slice(tmp, func(i, j int) bool {
				return i > j
			})
			Expect(tmp[0]).Should(BeNumerically("==", tmp[1]))
			Expect(len(tmp) == 2 || tmp[1] < tmp[2]).Should(BeTrue())

			actual, err := minimumIndex(seq)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(actual).Should(BeNumerically("==", 1))
		})
	})
})
