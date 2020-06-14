package solve

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Solve collect max diamond problem", func() {
	type testData struct {
		data     [][]int
		expected int
	}

	DescribeTable("문제를 풀었다.", func(d testData) {
		actual := collectMax(d.data)
		Expect(actual).Should(BeNumerically("==", d.expected))
	},
		Entry("test0", testData{[][]int{
			{0, 1, -1},
			{1, 0, -1},
			{1, 1, 1},
		}, 5}),
		Entry("test1", testData{[][]int{
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, 7}),
		Entry("test2", testData{[][]int{
			{0, 1, 1},
			{1, 0, -1},
			{1, 1, -1},
		}, 0}),
		Entry("test3_0", testData{[][]int{{}, {}}, 0}),
		Entry("test3_1", testData{[][]int{}, 0}),
	)

	It("combinations 함수 검증", func() {
		actual := combinations(3, 2)
		Expect(actual).Should(HaveLen(3))
		actual = combinations(4, 2)
		Expect(actual).Should(HaveLen(6))
		actual = combinations(4, 3)
		Expect(actual).Should(HaveLen(4))
	})

	It("find 함수 검증", func() {
		Expect(array([]int{1, 2, 3}).find(1)).Should(BeTrue())
		Expect(array([]int{1, 2, 3}).find(4)).Should(BeFalse())
		Expect(array([]int{}).find(1)).Should(BeFalse())
	})

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long string", func() {
			data := [][]int{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 2, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
			}

			Expect(collectMax(data)).Should(BeNumerically("==", 23))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 5), "시간 초과")
	}, 3)
})
