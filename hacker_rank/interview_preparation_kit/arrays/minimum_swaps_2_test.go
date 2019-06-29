package arrays

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/minimum-swaps-2/problem", func() {
	DescribeTable("문제를 풀었다",
		func(arr []int32, cnt int) {
			actual := minimumSwaps(arr)
			Expect(actual).Should(BeNumerically("==", cnt))
		},
		Entry("test_0", []int32{4, 3, 1, 2}, 3),
		Entry("test_1", []int32{2, 3, 4, 1, 5}, 3),
		Entry("test_2", []int32{1, 3, 5, 2, 4, 6, 7}, 3),
		Entry("test_3", []int32{7, 1, 3, 2, 4, 5, 6}, 5),
		Entry("test_4", []int32{1, 3, 4, 5, 6}, -1),
	)

	Context("find 함수는", func() {
		arr := []int32{1, 2, 3, 4, 5}
		length := len(arr)

		It("배열에서 주어진 오프셋 위치부터 타겟을 잘 찾는다.", func() {
			actual, ok := find(arr, int32(3), 2, length)
			Expect(ok).Should(BeTrue())
			Expect(actual).Should(Equal(2))
		})
		It("오프셋 앞쪽의 타겟은 찾지 않는다.", func() {
			_, ok := find(arr, int32(2), 2, length)
			Expect(ok).Should(BeFalse())
		})
	})
})
