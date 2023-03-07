package performance

import "testing"

type T struct {
	a *[N]int
}

var (
	globalT = T{a: &[N]int{}}
)

//go:noinline
func g0(t *T) {
	for i := range t.a {
		t.a[i] = i
	}
}

//go:noinline
func g1(t *T) {
	_ = *t.a
	for i := range t.a {
		t.a[i] = i
	}
}

//go:noinline
func g2(t *T) {
	a := t.a
	_ = *a
	for i := range a {
		a[i] = i
	}
}

//go:noinline
func g3(t *T) {
	a := t.a[:]
	for i := range a {
		a[i] = i
	}
}

func BenchmarkG0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g0(&globalT)
	}
}

func BenchmarkG1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g1(&globalT)
	}
}

func BenchmarkG2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g2(&globalT)
	}
}

func BenchmarkG3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g3(&globalT)
	}
}
