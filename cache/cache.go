package cache

import "math/big"

type Cache interface {
	Add(index uint64, result *big.Int) (err error)
	Clear() (err error)
	Get(index uint64) (result *big.Int, hit bool)
	GetRange(start, end uint64) (val []*big.Int, err error)
	Size() (size uint64, err error)
}

type NewCache func() Cache
