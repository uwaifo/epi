// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

import (
	"math/rand"
	"testing"
)

func TestIntStackMax(t *testing.T) {
	testStackInterface(t, new(IntStackMax), []interface{}{minInt, 0, maxInt})

	// Order of tests matter!
	tests := []struct {
		in   interface{}
		want interface{}
	}{
		{nil, nil}, // Test an max on empty stack.
		{4, 4},
		{3, 4},
		{1, 4},
		{4, 4},
		{5, 5},
		{5, 5},
		{0, 5},
		{9, 9},
		{6, 9},
		{8, 9},
		{7, 9},
	}
	s := new(IntStackMax)
	for _, test := range tests {
		if test.in != nil { // Don't push a nil value.
			s.Push(test.in)
		}
		if got := s.Max(); got != test.want {
			t.Errorf("%+v got s.Max() = %v; want %v", s, got, test.want)
		}
	}
	for i := len(tests) - 1; i >= 0; i-- {
		if got, want := s.Max(), tests[i].want; got != want {
			t.Errorf("%+v got s.Max() = %v; want %v", s, got, want)
		}
		s.Pop()
	}
}

func benchIntStackMax(b *testing.B, size int) {
	s := new(IntStackMax)
	max := -1
	for _, n := range rand.New(rand.NewSource(int64(size))).Perm(size) {
		s.Push(n)
		if n > max {
			max = n
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Max()
	}
}

func BenchmarkIntStackMax1e2(b *testing.B) { benchIntStackMax(b, 1e2) }
func BenchmarkIntStackMax1e4(b *testing.B) { benchIntStackMax(b, 1e4) }
func BenchmarkIntStackMax1e6(b *testing.B) { benchIntStackMax(b, 1e6) }
