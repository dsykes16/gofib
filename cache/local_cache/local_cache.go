package local_cache

import (
	"math/big"

	"github.com/dsykes16/gofib/cache"
)

type LocalCache struct {
	values map[uint64]*big.Int
}

func New() cache.Cache {
	return &LocalCache{values: makeCache()}
}

func (c *LocalCache) Add(index uint64, result *big.Int) (err error) {
	c.values[index] = result
	return nil
}

func (c *LocalCache) Clear() (err error) {
	c.values = makeCache()
	return nil
}

func (c *LocalCache) Get(index uint64) (result *big.Int, hit bool) {
	result, hit = c.values[index]
	return
}

func (c *LocalCache) GetRange(start, end uint64) (vals []*big.Int, err error) {
	for i := start; i < end; i++ {
		if n, hit := c.values[i]; hit {
			vals = append(vals, n)
		}
	}
	return
}

func (c *LocalCache) Size() (size uint64, err error) {
	return uint64(len(c.values)), nil
}

func makeCache() map[uint64]*big.Int {
	return make(map[uint64]*big.Int)
}
