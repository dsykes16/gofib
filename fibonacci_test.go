package gofib_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/dsykes16/gofib"
)

var _ = Describe("Fibonacci", func() {
	var fibonacci Fibonacci

	DescribeTable("Test basic Fib function to 10th Fibonacci number",
		func(index uint64, expected *big.Int) {
			Expect(fibonacci.Fib(index)).To(Equal(expected))
		},
		Entry("Fib(0)", uint64(0), big.NewInt(0)),
		Entry("Fib(1)", uint64(1), big.NewInt(1)),
		Entry("Fib(2)", uint64(2), big.NewInt(1)),
		Entry("Fib(3)", uint64(3), big.NewInt(2)),
		Entry("Fib(4)", uint64(4), big.NewInt(3)),
		Entry("Fib(5)", uint64(5), big.NewInt(5)),
		Entry("Fib(6)", uint64(6), big.NewInt(8)),
		Entry("Fib(7)", uint64(7), big.NewInt(13)),
		Entry("Fib(8)", uint64(8), big.NewInt(21)),
		Entry("Fib(9)", uint64(9), big.NewInt(34)),
		Entry("Fib(10)", uint64(10), big.NewInt(55)),
	)
})

var _ = Describe("LocalMemoizedFibonacci", func() {
	fibonacci := LocalMemoizedFibbonacci()

	DescribeTable("Test locally cached Fib function to 10th Fibonacci number",
		func(index uint64, expected *big.Int) {
			Expect(fibonacci.Fib(index)).To(Equal(expected))
		},
		Entry("Fib(0)", uint64(0), big.NewInt(0)),
		Entry("Fib(1)", uint64(1), big.NewInt(1)),
		Entry("Fib(2)", uint64(2), big.NewInt(1)),
		Entry("Fib(3)", uint64(3), big.NewInt(2)),
		Entry("Fib(4)", uint64(4), big.NewInt(3)),
		Entry("Fib(5)", uint64(5), big.NewInt(5)),
		Entry("Fib(6)", uint64(6), big.NewInt(8)),
		Entry("Fib(7)", uint64(7), big.NewInt(13)),
		Entry("Fib(8)", uint64(8), big.NewInt(21)),
		Entry("Fib(9)", uint64(9), big.NewInt(34)),
		Entry("Fib(10)", uint64(10), big.NewInt(55)),
	)
})
