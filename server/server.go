package server

import (
	"context"
	"errors"

	"github.com/dsykes16/gofib/fibonacci"
	pb "github.com/dsykes16/gofib/protos"
)

type server struct {
	fib *fibonacci.Fibonacci
}

func New(fibonacci *fibonacci.Fibonacci) *server {
	return &server{fib: fibonacci}
}

func (s *server) GetFib(ctx context.Context, request *pb.FibonacciRequest) (result *pb.FibonacciResult, err error) {
	if request == nil {
		err = errors.New("received nil request")
	} else {
		result = &pb.FibonacciResult{
			Index:  request.Index,
			Result: s.fib.Fib(request.Index).String(),
		}
	}
	return
}

func (s *server) GetFibRange(ctx context.Context, request *pb.FibonacciRangeRequest) (result *pb.FibonacciRangeResult, err error) {
	if request == nil {
		err = errors.New("received nil request")
	} else {
		result = &pb.FibonacciRangeResult{
			Result: []*pb.FibonacciResult{},
		}
		for i := request.Start; i < request.End; i++ {
			r := &pb.FibonacciResult{Index: i, Result: s.fib.Fib(i).String()}
			result.Result = append(result.Result, r)
		}
	}
	return
}

func (s *server) GetCacheSize(ctx context.Context, request *pb.Empty) (size *pb.Size, err error) {
	if request == nil {
		err = errors.New("received nil request")
	} else {
		size = &pb.Size{}
		size.Size, err = s.fib.Cache.Size()
	}
	return
}

func (s *server) GetCacheSizeForRange(ctx context.Context, request *pb.FibonacciRangeRequest) (result *pb.Size, err error) {
	if request == nil {
		return nil, errors.New("received nil request")
	}

	resr, err := s.fib.Cache.GetRange(request.Start, request.End)
	if err == nil {
		result = &pb.Size{Size: uint64(len(resr))}
	}
	return
}

func (s *server) ClearCache(ctx context.Context, request *pb.Empty) (result *pb.Empty, err error) {
	if request == nil {
		return nil, errors.New("received nil request")
	}
	result = &pb.Empty{}
	err = s.fib.Cache.Clear()
	return
}
