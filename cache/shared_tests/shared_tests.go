package shared_tests

import (
	"math/big"

	"github.com/dsykes16/gofib/cache"
	. "github.com/onsi/ginkgo" //lint:ignore ST1001 shared testing functions
	. "github.com/onsi/gomega" //lint:ignore ST1001 shared testing functions
)

func SharedCacheTests(newCache cache.NewCache) {
	Describe("Cache tests", func() {

		Context("initially", func() {
			cache := newCache()

			Specify("the cache size is 0", func() {
				size, err := cache.Size()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(size).Should(BeZero())
			})

			It("returns hit=false for non-existent value", func() {
				_, hit := cache.Get(uint64(0))
				Expect(hit).Should(BeFalse())
			})
		})

		Context("when a new item is added", func() {
			cache := newCache()
			cache.Add(uint64(5), big.NewInt(5))

			It("has 1 more item than previously", func() {
				Expect(cache.Size()).Should(Equal(uint64(1)))
			})

			It("returns the new value correctly by index", func() {
				val, hit := cache.Get(uint64(5))
				Expect(val).Should(Equal(big.NewInt(5)))
				Expect(hit).Should(BeTrue())
			})

			It("returns the value correctly given an inclusive index range", func() {
				vals, err := cache.GetRange(uint64(0), uint64(6))
				Expect(vals).Should(Equal([]*big.Int{big.NewInt(5)}))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when an existing item is added", func() {
			cache := newCache()
			cache.Add(uint64(5), big.NewInt(5))
			cache.Add(uint64(5), big.NewInt(5))

			It("has the same number of items", func() {
				size, err := cache.Size()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(size).Should(Equal(uint64(1)))
			})
		})

		Context("when the cache size is nonzero and it is cleared", func() {
			cache := newCache()
			cache.Add(uint64(5), big.NewInt(5))
			cache.Clear()

			It("should not contain previously cached value", func() {
				_, hit := cache.Get(uint64(5))
				Expect(hit).Should(BeFalse())
			})

			Specify("cache size is 0", func() {
				size, err := cache.Size()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(size).Should(BeZero())
			})
		})

	})
}
