package tutorial30daysofcode

import (
	"time"

	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem", func() {
	type testData struct {
		from, to string
		result   int32
	}

	const layout = "2 1 2006"

	DescribeTable("문제를 풀었다",
		func(d testData) {
			from, err := time.Parse(layout, convertCalendar(d.from))
			Expect(err).ShouldNot(HaveOccurred())
			to, err := time.Parse(layout, convertCalendar(d.to))
			Expect(err).ShouldNot(HaveOccurred())

			Expect(nestedLogic(from, to)).Should(BeNumerically("==", d.result))
		},
		Entry("0", testData{"1 1 1", "8 8 8", 0}),
		Entry("0", testData{"6 6 2015", "9 6 2015", 0}),
		Entry("0", testData{"9 6 2015", "6 6 2015", 45}),
		Entry("0", testData{"6 7 2015", "6 6 2015", 500}),
		Entry("0", testData{"6 6 2016", "6 6 2015", 10000}),
	)
})
