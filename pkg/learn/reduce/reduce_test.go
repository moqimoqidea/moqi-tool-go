// Package reduce
// @author moqi
// On 2021/9/5 23:38:05
package reduce

import (
	"robpike.io/filter"
	"testing"
)

func mul(a, b int) int {
	return a * b
}

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}

// Just Use https://github.com/robpike/filter/blob/master/reduce_test.go
func TestReduce(t *testing.T) {
	const size = 10
	a := make([]int, size)
	for i := range a {
		a[i] = i + 1
	}
	for i := 1; i < 10; i++ {
		out := filter.Reduce(a[:i], mul, 1).(int)
		expect := fac(i)
		if expect != out {
			t.Fatalf("expected %d got %d", expect, out)
		}
	}
}
