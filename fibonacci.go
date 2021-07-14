package gofib

type Fibonacci struct{}

func (f Fibonacci) Fib(index uint) (result uint) {
	if index < 2 {
		return index
	}
	return f.Fib(index-1) + f.Fib(index-2)
}
