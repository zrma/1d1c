package lv2hard

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/median-of-two-sorted-arrays/", func() {
	type testData struct {
		arr1, arr2 []int
		expected   float64
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		actual := findMedianSortedArrays(data.arr1, data.arr2)
		Expect(actual).Should(Equal(data.expected))
	},
		Entry("0", testData{
			arr1:     []int{1, 3},
			arr2:     []int{2},
			expected: 2.0,
		}),
		Entry("1", testData{
			arr1:     []int{1, 2},
			arr2:     []int{3, 4},
			expected: 2.5,
		}),
		Entry("2", testData{
			arr1:     []int{1, 2},
			arr2:     []int{100},
			expected: 2.0,
		}),
		Entry("3-0", testData{
			arr1:     []int{1, 200},
			arr2:     []int{10, 90},
			expected: 50.0,
		}),
		Entry("3-1", testData{
			arr1:     []int{1, 10, 200},
			arr2:     []int{90},
			expected: 50.0,
		}),
		Entry("4-0", testData{
			arr1:     []int{1, 2, 2, 3},
			arr2:     []int{100},
			expected: 2.0,
		}),
		Entry("4-1", testData{
			arr1:     []int{1, 2, 2, 3, 3, 3},
			arr2:     []int{100},
			expected: 3.0,
		}),
		Entry("5", testData{
			arr1:     []int{2, 2, 3, 3, 3, 100},
			arr2:     []int{1},
			expected: 3.0,
		}),
		Entry("5-1", testData{
			arr1:     []int{2, 2, 2, 2, 2, 2},
			arr2:     []int{2},
			expected: 2.0,
		}),
		Entry("5-2", testData{
			arr1:     []int{2},
			arr2:     []int{2, 2, 2, 2, 2, 2},
			expected: 2.0,
		}),
	)

	It("mergeSortedArray", func() {
		actual := mergeSortedArray([]int{1, 3, 4, 8}, []int{2, 4, 5, 7})
		Expect(actual).Should(Equal([]int{1, 2, 3, 4, 4, 5, 7, 8}))
	})
})
