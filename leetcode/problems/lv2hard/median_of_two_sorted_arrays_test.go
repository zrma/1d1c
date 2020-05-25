package lv2hard

import (
	"math"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/median-of-two-sorted-arrays/", func() {
	Context("함수 전체 테스트", func() {
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

			By("By bin search solution", func() {
				actual := findMedianSortedArraysWithBinSearchSolution(data.arr1, data.arr2)
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
			Entry("6-1", testData{
				arr1:     []int{3, 4},
				arr2:     []int{1, 2},
				expected: 2.5,
			}),
			Entry("6-2", testData{
				arr1:     []int{1},
				arr2:     []int{1},
				expected: 1,
			}),
			Entry("6-2", testData{
				arr1:     []int{1},
				arr2:     []int{2, 3, 4},
				expected: 2.5,
			}),
		)
	})

	Context("유틸 함수 테스트", func() {
		nums := arr([]int{0, 1, 2, 3, 4})

		type testData struct {
			idx      int
			expected int
		}

		DescribeTable("arr.getVal", func(data testData) {
			actual := arr(nums).getVal(data.idx)
			Expect(actual).Should(Equal(data.expected))
		},
			Entry("0", testData{idx: -1, expected: math.MinInt32}),
			Entry("1", testData{idx: 0, expected: 0}),
			Entry("2", testData{idx: 4, expected: 4}),
			Entry("3", testData{idx: 5, expected: math.MaxInt32}),
		)
	})
})

//goos: linux
//goarch: amd64
//pkg: github.com/zrma/1d1c/leetcode/problems/lv2hard
//BenchmarkFindMedian1/By_merge-16                         127035747               9.02   ns/op
//BenchmarkFindMedian1/By_index-16                        1000000000               0.929  ns/op
//BenchmarkFindMedian1/By_bin_search-16                   1000000000               1.12   ns/op
//BenchmarkFindMedian1/By_bin_search_solution-16          1000000000               0.803  ns/op
//BenchmarkFindMedian2/By_merge-16                          75473596              13.3    ns/op
//BenchmarkFindMedian2/By_index-16                        1000000000               0.863  ns/op
//BenchmarkFindMedian2/By_bin_search-16                   1000000000               0.850  ns/op
//BenchmarkFindMedian2/By_bin_search_solution-16          1000000000               0.682  ns/op
//BenchmarkFindMedian4/By_merge-16                          46541391              23.9    ns/op
//BenchmarkFindMedian4/By_index-16                         908933222               1.33   ns/op
//BenchmarkFindMedian4/By_bin_search-16                    723928178               1.67   ns/op
//BenchmarkFindMedian4/By_bin_search_solution-16           833685333               1.43   ns/op
//BenchmarkFindMedian8/By_merge-16                          27980440              40.8    ns/op
//BenchmarkFindMedian8/By_index-16                         658515704               1.81   ns/op
//BenchmarkFindMedian8/By_bin_search-16                   1000000000               1.14   ns/op
//BenchmarkFindMedian8/By_bin_search_solution-16           922768860               1.31   ns/op
//BenchmarkFindMedian16/By_merge-16                         21129810              55.5    ns/op
//BenchmarkFindMedian16/By_index-16                        493847685               2.43   ns/op
//BenchmarkFindMedian16/By_bin_search-16                   756249932               1.61   ns/op
//BenchmarkFindMedian16/By_bin_search_solution-16          867374154               1.40   ns/op
//BenchmarkFindMedian32/By_merge-16                         17221905              67.9    ns/op
//BenchmarkFindMedian32/By_index-16                        587363749               2.07   ns/op
//BenchmarkFindMedian32/By_bin_search-16                  1000000000               1.10   ns/op
//BenchmarkFindMedian32/By_bin_search_solution-16         1000000000               1.11   ns/op
//BenchmarkFindMedian64/By_merge-16                          6141870             201      ns/op
//BenchmarkFindMedian64/By_index-16                        123069411               9.83   ns/op
//BenchmarkFindMedian64/By_bin_search-16                  1000000000               1.11   ns/op
//BenchmarkFindMedian64/By_bin_search_solution-16          919578985               1.27   ns/op
//BenchmarkFindMedian128/By_merge-16                         3690441             325      ns/op
//BenchmarkFindMedian128/By_index-16                        80059777              14.6    ns/op
//BenchmarkFindMedian128/By_bin_search-16                  638245250               1.86   ns/op
//BenchmarkFindMedian128/By_bin_search_solution-16         617231878               1.94   ns/op
//BenchmarkFindMedian256/By_merge-16                         2252612             503      ns/op
//BenchmarkFindMedian256/By_index-16                        67390363              18.0    ns/op
//BenchmarkFindMedian256/By_bin_search-16                  755317909               1.59   ns/op
//BenchmarkFindMedian256/By_bin_search_solution-16         705693475               1.71   ns/op
//BenchmarkFindMedian512/By_merge-16                         1644506             732      ns/op
//BenchmarkFindMedian512/By_index-16                        19918366              61.5    ns/op
//BenchmarkFindMedian512/By_bin_search-16                 1000000000               1.11   ns/op
//BenchmarkFindMedian512/By_bin_search_solution-16         914383945               1.29   ns/op
//BenchmarkFindMedian1024/By_merge-16                        1335741             891      ns/op
//BenchmarkFindMedian1024/By_index-16                       26047544              45.3    ns/op
//BenchmarkFindMedian1024/By_bin_search-16                 636004533               1.85   ns/op
//BenchmarkFindMedian1024/By_bin_search_solution-16        853142835               1.39   ns/op
//BenchmarkFindMedian2048/By_merge-16                         479786            2138      ns/op
//BenchmarkFindMedian2048/By_index-16                        7599582             161      ns/op
//BenchmarkFindMedian2048/By_bin_search-16                1000000000               0.816  ns/op
//BenchmarkFindMedian2048/By_bin_search_solution-16       1000000000               0.977  ns/op
//BenchmarkFindMedian4096/By_merge-16                         186649            6579      ns/op
//BenchmarkFindMedian4096/By_index-16                        2550916             471      ns/op
//BenchmarkFindMedian4096/By_bin_search-16                 299504170               4.01   ns/op
//BenchmarkFindMedian4096/By_bin_search_solution-16        493479288               2.43   ns/op
//BenchmarkFindMedian8192/By_merge-16                         122881            9501      ns/op
//BenchmarkFindMedian8192/By_index-16                        1304148             924      ns/op
//BenchmarkFindMedian8192/By_bin_search-16                1000000000               1.11   ns/op
//BenchmarkFindMedian8192/By_bin_search_solution-16        947535735               1.28   ns/op
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

	b.Run("By bin search solution", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				findMedianSortedArraysWithBinSearchSolution(nums1, nums2)
			}
		})
	})
}
