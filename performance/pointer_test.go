package performance

import "testing"

const N = 1000

var a [N]int

//go:noinline
func f0(a *[N]int) {
	for i := range a {
		a[i] = i
	}
}

//go:noinline
func f1(a *[N]int) {
	_ = *a
	for i := range a {
		a[i] = i
	}
}

func BenchmarkF0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f0(&a)
	}
}

func BenchmarkF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1(&a)
	}
}
