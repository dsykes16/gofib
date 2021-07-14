package gofib

import (
	"math/big"

	"github.com/dsykes16/gofib/cache"
	"github.com/dsykes16/gofib/cache/local_cache"
)

type Fibonacci struct {
	cache cache.Cache
}

func (f Fibonacci) Fib(index uint64) (result *big.Int) {
	if index < 2 {
		return big.NewInt(int64(index))
	}

	if f.cache != nil {
		if result, hit := f.cache.Get(index); hit {
			return result
		}
	}

	result = big.NewInt(0) // allocate memory for new big.Int
	result = result.Add(f.Fib(index-1), f.Fib(index-2))

	if f.cache != nil {
		f.cache.Add(index, result)
	}

	return
}

func LocalMemoizedFibbonacci() (fibonacci *Fibonacci) {
	return &Fibonacci{cache: local_cache.NewLocalCache()}
}
