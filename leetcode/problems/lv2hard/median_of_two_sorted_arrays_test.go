package lv2hard

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log("https://leetcode.com/problems/median-of-two-sorted-arrays/")

	for i, tt := range []struct {
		arr1, arr2 []int
		want       float64
	}{
		{
			arr1: []int{1, 3},
			arr2: []int{2},
			want: 2.0,
		},
		{
			arr1: []int{1, 2},
			arr2: []int{3, 4},
			want: 2.5,
		},
		{
			arr1: []int{1, 2},
			arr2: []int{100},
			want: 2.0,
		},
		{
			arr1: []int{1, 200},
			arr2: []int{10, 90},
			want: 50.0,
		},
		{
			arr1: []int{1, 10, 200},
			arr2: []int{90},
			want: 50.0,
		},
		{
			arr1: []int{1, 2, 2, 3},
			arr2: []int{100},
			want: 2.0,
		},
		{
			arr1: []int{1, 2, 2, 3, 3, 3},
			arr2: []int{100},
			want: 3.0,
		},
		{
			arr1: []int{2, 2, 3, 3, 3, 100},
			arr2: []int{1},
			want: 3.0,
		},
		{
			arr1: []int{2, 2, 2, 2, 2, 2},
			arr2: []int{2},
			want: 2.0,
		},
		{
			arr1: []int{2},
			arr2: []int{2, 2, 2, 2, 2, 2},
			want: 2.0,
		},
		{
			arr1: []int{2, 3, 4, 4, 4, 4, 4, 4, 4},
			arr2: []int{1, 5, 5, 5},
			want: 4,
		},
		{
			arr1: []int{3, 4},
			arr2: []int{1, 2},
			want: 2.5,
		},
		{
			arr1: []int{1},
			arr2: []int{1},
			want: 1,
		},
		{
			arr1: []int{1},
			arr2: []int{2, 3, 4},
			want: 2.5,
		},
	} {
		t.Run(fmt.Sprintf("%d by merge", i), func(t *testing.T) {
			got := findMedianSortedArraysWithMerge(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d by index", i), func(t *testing.T) {
			got := findMedianSortedArrays(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d by bin search", i), func(t *testing.T) {
			got := findMedianSortedArraysWithBinSearch(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d by bin search solution\"", i), func(t *testing.T) {
			got := findMedianSortedArraysWithBinSearchSolution(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetVal(t *testing.T) {
	for _, tt := range []struct {
		given int
		want  int
	}{
		{-1, math.MinInt32},
		{0, 0},
		{4, 4},
		{5, math.MaxInt32},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			nums := arr([]int{0, 1, 2, 3, 4})
			got := nums.getVal(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

//goos: linux
//goarch: amd64
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
