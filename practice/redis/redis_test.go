package redis

import (
	"github.com/alicebob/miniredis/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redis 테스트", func() {
	var mockRedis *miniredis.Miniredis

	BeforeEach(func() {
		var err error
		mockRedis, err = miniredis.Run()
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		mockRedis.Close()
	})

	It("HGet, HSet", func() {
		const (
			hash = "hash123"
			key  = "key1"
			val  = "val2"
		)

		HSet(mockRedis, hash, key, val)
		actual := HGet(mockRedis, hash, key)
		Expect(actual).Should(Equal(val))
	})
})
