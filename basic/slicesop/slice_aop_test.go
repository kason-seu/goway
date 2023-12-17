package slicesop

import (
	"testing"
)

func Test_slicesop_fn(t *testing.T) {
	slicesop_fn()
}

const turns = 1000000

func Benchmark_slicesop_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < turns; j++ {
			s = append(s, j)
		}
		//fmt.Println("s", s, "len :", len(s), "cap : ", cap(s))
	}
}
 
func Benchmark_slicesop_fn_prebuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int = make([]int, turns, turns)
		for j := 0; j < turns; j++ {
			s = append(s, j)
		}
		//fmt.Println("s", s, "len :", len(s), "cap : ", cap(s))
	}
}