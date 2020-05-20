package lv2hard

import (
	"math/rand"
	"testing"

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
		By("By merge", func() {
			actual := findMedianSortedArraysWithMerge(data.arr1, data.arr2)
			Expect(actual).Should(Equal(data.expected))
		})

		By("By index", func() {
			actual := findMedianSortedArrays(data.arr1, data.arr2)
			Expect(actual).Should(Equal(data.expected))
		})

		By("By bin search", func() {
			actual := findMedianSortedArraysWithBinSearch(data.arr1, data.arr2)
			Expect(actual).Should(Equal(data.expected))
		})
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
		Entry("6", testData{
			arr1:     []int{2, 3, 4, 4, 4, 4, 4, 4, 4},
			arr2:     []int{1, 5, 5, 5},
			expected: 4,
		}),
	)
})

//goos: linux
//goarch: amd64
//BenchmarkFindMedian1/By_merge-16                126576813                9.36 ns/op
//BenchmarkFindMedian1/By_index-16                1000000000               0.864 ns/op
//BenchmarkFindMedian1/By_bin_search-16           1000000000               0.734 ns/op
//BenchmarkFindMedian2/By_merge-16                86559476                12.5 ns/op
//BenchmarkFindMedian2/By_index-16                1000000000               0.892 ns/op
//BenchmarkFindMedian2/By_bin_search-16           1000000000               0.609 ns/op
//BenchmarkFindMedian4/By_merge-16                51975051                23.1 ns/op
//BenchmarkFindMedian4/By_index-16                946856868                1.24 ns/op
//BenchmarkFindMedian4/By_bin_search-16           849017156                1.40 ns/op
//BenchmarkFindMedian8/By_merge-16                26094500                42.2 ns/op
//BenchmarkFindMedian8/By_index-16                672385303                1.83 ns/op
//BenchmarkFindMedian8/By_bin_search-16           925608818                1.30 ns/op
//BenchmarkFindMedian16/By_merge-16               21669801                52.5 ns/op
//BenchmarkFindMedian16/By_index-16               494116309                2.50 ns/op
//BenchmarkFindMedian16/By_bin_search-16          882739788                1.39 ns/op
//BenchmarkFindMedian32/By_merge-16               18645150                64.3 ns/op
//BenchmarkFindMedian32/By_index-16               571174806                2.19 ns/op
//BenchmarkFindMedian32/By_bin_search-16          1000000000               1.09 ns/op
//BenchmarkFindMedian64/By_merge-16                6408602               194 ns/op
//BenchmarkFindMedian64/By_index-16               122519014                9.77 ns/op
//BenchmarkFindMedian64/By_bin_search-16          918924092                1.28 ns/op
//BenchmarkFindMedian128/By_merge-16               3674931               307 ns/op
//BenchmarkFindMedian128/By_index-16              81134256                14.8 ns/op
//BenchmarkFindMedian128/By_bin_search-16         585064473                1.99 ns/op
//BenchmarkFindMedian256/By_merge-16               2416862               503 ns/op
//BenchmarkFindMedian256/By_index-16              62001084                18.2 ns/op
//BenchmarkFindMedian256/By_bin_search-16         690182328                1.73 ns/op
//BenchmarkFindMedian512/By_merge-16               1669188               728 ns/op
//BenchmarkFindMedian512/By_index-16              18578439                63.8 ns/op
//BenchmarkFindMedian512/By_bin_search-16         953511543                1.30 ns/op
//BenchmarkFindMedian1024/By_merge-16              1357064               888 ns/op
//BenchmarkFindMedian1024/By_index-16             26570716                46.7 ns/op
//BenchmarkFindMedian1024/By_bin_search-16        811126491                1.48 ns/op
//BenchmarkFindMedian2048/By_merge-16               507120              2079 ns/op
//BenchmarkFindMedian2048/By_index-16              7414038               164 ns/op
//BenchmarkFindMedian2048/By_bin_search-16        1000000000               1.05 ns/op
//BenchmarkFindMedian4096/By_merge-16               179924              6438 ns/op
//BenchmarkFindMedian4096/By_index-16              2484904               485 ns/op
//BenchmarkFindMedian4096/By_bin_search-16        467696221                2.52 ns/op
//BenchmarkFindMedian8192/By_merge-16               122500              9718 ns/op
//BenchmarkFindMedian8192/By_index-16              1262004               953 ns/op
//BenchmarkFindMedian8192/By_bin_search-16        940063140                1.31 ns/op
func BenchmarkFindMedian1(b *testing.B)    { benchmarkFindMedianSortedArrays(1, b) }
func BenchmarkFindMedian2(b *testing.B)    { benchmarkFindMedianSortedArrays(2, b) }
func BenchmarkFindMedian4(b *testing.B)    { benchmarkFindMedianSortedArrays(4, b) }
func BenchmarkFindMedian8(b *testing.B)    { benchmarkFindMedianSortedArrays(8, b) }
func BenchmarkFindMedian16(b *testing.B)   { benchmarkFindMedianSortedArrays(16, b) }
func BenchmarkFindMedian32(b *testing.B)   { benchmarkFindMedianSortedArrays(32, b) }
func BenchmarkFindMedian64(b *testing.B)   { benchmarkFindMedianSortedArrays(64, b) }
func BenchmarkFindMedian128(b *testing.B)  { benchmarkFindMedianSortedArrays(128, b) }
func BenchmarkFindMedian256(b *testing.B)  { benchmarkFindMedianSortedArrays(256, b) }
func BenchmarkFindMedian512(b *testing.B)  { benchmarkFindMedianSortedArrays(512, b) }
func BenchmarkFindMedian1024(b *testing.B) { benchmarkFindMedianSortedArrays(1024, b) }
func BenchmarkFindMedian2048(b *testing.B) { benchmarkFindMedianSortedArrays(2048, b) }
func BenchmarkFindMedian4096(b *testing.B) { benchmarkFindMedianSortedArrays(4096, b) }
func BenchmarkFindMedian8192(b *testing.B) { benchmarkFindMedianSortedArrays(8192, b) }

func benchmarkFindMedianSortedArrays(cnt1 int, b *testing.B) {
	b.StopTimer()
	nums1 := make([]int, cnt1)
	for i := 0; i < cnt1; i++ {
		nums1[i] = rand.Intn(cnt1)
	}

	cnt2 := rand.Intn(cnt1) + 1
	nums2 := make([]int, cnt2)
	for i := 0; i < cnt2; i++ {
		nums2[i] = rand.Intn(cnt1)
	}
	b.StartTimer()

	b.Run("By merge", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				findMedianSortedArraysWithMerge(nums1, nums2)
			}
		})
	})

	b.Run("By index", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				findMedianSortedArrays(nums1, nums2)
			}
		})
	})

	b.Run("By bin search", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				findMedianSortedArraysWithBinSearch(nums1, nums2)
			}
		})
	})
}
