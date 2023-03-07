package performance

import "testing"

var s = make([]int, 1024)
var r int

//go:noinline
func h0(sum *int, s []int) {
	for _, v := range s {
		*sum += v
	}
}

//go:noinline
func h1(sum *int, s []int) {
	var n = *sum
	for _, v := range s {
		n += v
	}
	*sum = n
}

func BenchmarkH0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h0(&r, s)
	}
}

func BenchmarkH1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h1(&r, s)
	}
}
