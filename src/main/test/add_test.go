// filename: add_test.go

package test

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd1(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestAdd2(t *testing.T) {
	if Add(2, 3) != 6 {
		//t.Fatal("result is wrong!")
	} else {
		t.Log("result is right")

	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 5)
	}
}
