package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			var arr []interface{}
			for _, data := range []int{1, 2, 3} {
				arr = append(arr, data)
			}
			printArray(arr...)

			arr = arr[:0]
			for _, data := range []string{"Hello", "World"} {
				arr = append(arr, data)
			}
			printArray(arr...)
		}, []string{
			"1",
			"2",
			"3",
			"Hello",
			"World",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
