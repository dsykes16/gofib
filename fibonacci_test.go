package rpcfib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/dsykes16/rpcfib"
)

var _ = Describe("Fibonacci", func() {
	var fibonacci Fibonacci

	DescribeTable("Test basic Fib function to 10th Fibonacci number",
		func(index uint, expected uint) {
			Expect(fibonacci.Fib(index)).To(Equal(expected))
		},
		Entry("Fib(0)", uint(0), uint(0)),
		Entry("Fib(1)", uint(1), uint(1)),
		Entry("Fib(2)", uint(2), uint(1)),
		Entry("Fib(3)", uint(3), uint(2)),
		Entry("Fib(4)", uint(4), uint(3)),
		Entry("Fib(5)", uint(5), uint(5)),
		Entry("Fib(6)", uint(6), uint(8)),
		Entry("Fib(7)", uint(7), uint(13)),
		Entry("Fib(8)", uint(8), uint(21)),
		Entry("Fib(9)", uint(9), uint(34)),
		Entry("Fib(10)", uint(10), uint(55)),
	)
})
