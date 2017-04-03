package iter

import (
	"testing"
)

// const loops = 1e6
const loops = 10

func BenchmarkForClause(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = 0; j < loops; j++ {
			j = j
		}
	}
	_ = j
}

func BenchmarkRangeIter(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = range N(loops) {
			j = j
		}
	}
	_ = j
}
