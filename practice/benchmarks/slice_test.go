package benchmarks

import (
	"testing"
)

func benchmarkSliceConstruction(i int64, b *testing.B) {
	b.Run("make", func(b *testing.B) {
		benchmarkMakeSlice(i, b)
	})

	b.Run("append", func(b *testing.B) {
		benchmarkAppendSlice(i, b)
	})
}

func benchmarkMakeSlice(i int64, b *testing.B) {
	cnt := i * int64(b.N)
	var arr []int64
	for i = 0; i < cnt; i++ {
		arr = append(arr, i)
	}
	_ = arr
}

func benchmarkAppendSlice(i int64, b *testing.B) {
	cnt := i * int64(b.N)
	arr := make([]int64, cnt)
	for i = 0; i < cnt; i++ {
		arr[i] = i
	}
	_ = arr
}

func BenchmarkSliceConstruction1(b *testing.B)    { benchmarkSliceConstruction(1, b) }
func BenchmarkSliceConstruction2(b *testing.B)    { benchmarkSliceConstruction(2, b) }
func BenchmarkSliceConstruction4(b *testing.B)    { benchmarkSliceConstruction(4, b) }
func BenchmarkSliceConstruction8(b *testing.B)    { benchmarkSliceConstruction(8, b) }
func BenchmarkSliceConstruction16(b *testing.B)   { benchmarkSliceConstruction(16, b) }
func BenchmarkSliceConstruction32(b *testing.B)   { benchmarkSliceConstruction(32, b) }
func BenchmarkSliceConstruction64(b *testing.B)   { benchmarkSliceConstruction(64, b) }
func BenchmarkSliceConstruction128(b *testing.B)  { benchmarkSliceConstruction(128, b) }
func BenchmarkSliceConstruction256(b *testing.B)  { benchmarkSliceConstruction(256, b) }
func BenchmarkSliceConstruction512(b *testing.B)  { benchmarkSliceConstruction(512, b) }
func BenchmarkSliceConstruction1024(b *testing.B) { benchmarkSliceConstruction(1024, b) }
