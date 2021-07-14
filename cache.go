package rpcfib

type Cache interface {
	Add(index, result uint) (err error)
	Clear() (err error)
	Get(index, result uint) (val uint, err error)
	GetRange(start, end uint) (val []uint, err error)
	Size() (size uint, err error)
}
