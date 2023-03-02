package pool

type Pool struct {
	Max  uint64
	Chan chan uint64
	fn   func(interface{})
}

func NewPool(max uint64, fn func(interface{})) *Pool {
	return &Pool{
		Max:  max,
		Chan: make(chan uint64, max),
		fn:   fn,
	}
}
func New(max uint64) *Pool {
	return &Pool{
		Max:  max,
		Chan: make(chan uint64, max),
	}
}
func (pool *Pool) Submit(fn func()) {
	pool.Chan <- 1
	go func() {
		fn()
		<-pool.Chan
	}()
}
func (pool *Pool) SubmitWithParam(param interface{}) {
	pool.Chan <- 1
	go func() {
		pool.fn(param)
		<-pool.Chan
	}()
}
func (pool *Pool) Close() {
	close(pool.Chan)
}
