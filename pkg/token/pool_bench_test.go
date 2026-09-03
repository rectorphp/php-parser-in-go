package token

import (
	"testing"
)

const amount = 100000

func BenchmarkPlain(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)

		for range amount {
			tokens = append(tokens, &Token{})
		}
	}
}

func BenchmarkSlice128(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 128)

		for range amount {
			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkSlice512(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 512)

		for range amount {
			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkSlice1024(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 1024)

		for range amount {
			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkSlice2048(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 2048)

		for range amount {
			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend128(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 128)

		for range amount {
			if len(slc) == 128 {
				slc = make([]Token, 0, 128)
			}

			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend512(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 512)

		for range amount {
			if len(slc) == 512 {
				slc = make([]Token, 0, 512)
			}

			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend1024(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 1024)

		for range amount {
			if len(slc) == 1024 {
				slc = make([]Token, 0, 1024)
			}

			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend2048(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		tokens := make([]*Token, 0, amount)
		slc := make([]Token, 0, 2048)

		for range amount {
			if len(slc) == 2048 {
				slc = make([]Token, 0, 2048)
			}

			slc = append(slc, Token{})
			tokens = append(tokens, &slc[len(slc)-1])
		}
	}
}

func BenchmarkPool128(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		pool := NewPool(128)
		tokens := make([]*Token, 0, amount)

		for range amount {
			tokens = append(tokens, pool.Get())
		}
	}
}

func BenchmarkPool512(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		pool := NewPool(512)
		tokens := make([]*Token, 0, amount)

		for range amount {
			tokens = append(tokens, pool.Get())
		}
	}
}

func BenchmarkPool1024(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		pool := NewPool(1024)
		tokens := make([]*Token, 0, amount)

		for range amount {
			tokens = append(tokens, pool.Get())
		}
	}
}

func BenchmarkPool2048(benchmark *testing.B) {
	for number := 0; number < benchmark.N; number++ {
		pool := NewPool(2048)
		tokens := make([]*Token, 0, amount)

		for range amount {
			tokens = append(tokens, pool.Get())
		}
	}
}
