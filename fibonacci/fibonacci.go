package fibonacci

import (
	"math/big"

	"github.com/dsykes16/gofib/cache"
)

type Fibonacci struct {
	Cache cache.Cache
}

func New(c cache.Cache) (fibonacci *Fibonacci) {
	fibonacci = &Fibonacci{
		Cache: c,
	}
	return
}

func (f Fibonacci) Fib(index uint64) (result *big.Int) {
	if index < 2 {
		return big.NewInt(int64(index))
	}

	if f.Cache != nil {
		if result, hit := f.Cache.Get(index); hit {
			return result
		}
	}

	result = big.NewInt(0) // allocate memory for new big.Int
	result = result.Add(f.Fib(index-1), f.Fib(index-2))

	if f.Cache != nil {
		f.Cache.Add(index, result)
	}

	return
}
