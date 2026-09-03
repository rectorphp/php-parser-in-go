package token

const DefaultBlockSize = 1024

type Pool struct {
	block []Token
	off   int
}

func NewPool(blockSize int) *Pool {
	return &Pool{
		block: make([]Token, blockSize),
	}
}

func (pool *Pool) Get() *Token {
	if len(pool.block) == 0 {
		return nil
	}

	if len(pool.block) == pool.off {
		pool.block = make([]Token, len(pool.block))
		pool.off = 0
	}

	pool.off++

	return &pool.block[pool.off-1]
}
