package position

const DefaultBlockSize = 1024

type Pool struct {
	block []Position
	off   int
}

func NewPool(blockSize int) *Pool {
	return &Pool{
		block: make([]Position, blockSize),
	}
}

func (pool *Pool) Get() *Position {
	if len(pool.block) == 0 {
		return nil
	}

	if len(pool.block) == pool.off {
		pool.block = make([]Position, len(pool.block))
		pool.off = 0
	}

	pool.off++

	return &pool.block[pool.off-1]
}
