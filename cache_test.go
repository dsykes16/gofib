package rpcfib_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Cache", func() {
	Context("initially", func() {
		It("has 0 items", func() {})
		Specify("the cache size is 0", func() {})
	})

	Context("when a new item is added", func() {
		It("has 1 more item than previously", func() {})
		It("returns the new value correctly by index", func() {})
		It("returns the value correctly given an inclusive index range", func() {})
		Specify("cache size increases by 1", func() {})
	})

	Context("when an existing item is added", func() {
		It("has the same number of items", func() {})
		Specify("cache size remains unchanged", func() {})
	})

	Context("when the cache size is nonzero and it is cleared", func() {
		It("should clear all entries", func() {})
		Specify("cache size is 0", func() {})
	})

})
