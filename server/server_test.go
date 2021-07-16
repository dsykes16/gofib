package server_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/dsykes16/gofib/fibonacci"
	pb "github.com/dsykes16/gofib/protos"
	"github.com/dsykes16/gofib/server"
)

// TODO: Create mocks for Fibonacci and Cache for properly isolated unit tests

var _ = Describe("gRPC Server Tests", func() {
	ctx := context.Background()
	fib := fibonacci.LocalMemoizedFibbonacci()
	s := server.New(fib)
	DescribeTable("GetFib",
		func(request *pb.FibonacciRequest, expectedResult *pb.FibonacciResult, expectError bool) {
			res, err := s.GetFib(ctx, request)
			if !expectError {
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(expectedResult))
			} else {
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			}
		},
		Entry(
			"valid request",
			&pb.FibonacciRequest{
				Index: uint64(5),
			},
			&pb.FibonacciResult{
				Index:  uint64(5),
				Result: "5",
			},
			false,
		),
		Entry(
			"invalid request",
			nil,
			nil,
			true,
		),
	)

	DescribeTable("GetFibRange",
		func(request *pb.FibonacciRangeRequest, expectedResult *pb.FibonacciRangeResult, expectError bool) {
			res, err := s.GetFibRange(ctx, request)
			if !expectError {
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(expectedResult))
			} else {
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			}
		},
		Entry(
			"valid request",
			&pb.FibonacciRangeRequest{
				Start: uint64(1),
				End:   uint64(4),
			},
			&pb.FibonacciRangeResult{
				Result: []*pb.FibonacciResult{
					{
						Index:  uint64(1),
						Result: "1",
					},
					{
						Index:  uint64(2),
						Result: "1",
					},
					{
						Index:  uint64(3),
						Result: "2",
					},
				},
			},
			false,
		),
		Entry(
			"invalid request",
			nil,
			nil,
			true,
		),
	)

	DescribeTable("GetCacheSize",
		func(request *pb.Empty, expectError bool) {
			res, err := s.GetCacheSize(ctx, request)
			if !expectError {
				Expect(err).NotTo(HaveOccurred())
				cacheSize, _ := fib.Cache.Size()
				expectedResult := &pb.Size{Size: cacheSize}
				Expect(res).To(Equal(expectedResult))
			} else {
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			}
		},
		Entry(
			"valid request",
			&pb.Empty{},
			false,
		),
		Entry(
			"invalid request",
			nil,
			true,
		),
	)

	DescribeTable("GetCacheSizeForRange",
		func(request *pb.FibonacciRangeRequest, expectedResult *pb.Size, expectError bool) {
			res, err := s.GetCacheSizeForRange(ctx, request)
			if !expectError {
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(expectedResult))
			} else {
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			}
		},
		Entry(
			"valid request",
			&pb.FibonacciRangeRequest{
				Start: uint64(0),
				End:   uint64(0),
			},
			&pb.Size{
				Size: uint64(0),
			},
			false,
		),
		Entry(
			"invalid request",
			nil,
			nil,
			true,
		),
	)

	Describe("ClearCache", func() {
		It("returns err given invalid request", func() {
			res, err := s.ClearCache(ctx, nil)
			Expect(res).To(BeNil())
			Expect(err).To(HaveOccurred())
		})
		It("clears the cache", func() {
			res, err := s.ClearCache(ctx, &pb.Empty{})
			Expect(res).To(Equal(&pb.Empty{}))
			Expect(err).NotTo(HaveOccurred())
			cacheSize, _ := fib.Cache.Size()
			Expect(cacheSize).To(Equal(uint64(0)))
		})
	})
})
