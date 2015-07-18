package main

import (
	"strings"
	"testing"
)

func TestTis100(t *testing.T) {
	m := NodeMap{
		{T21Node},
		{T21Node},
	}

	in := make(chan int)

	tis := NewTis100(m)
	tis.Input(in, 0)
	out := tis.Output(1)

	go func() {
		for i := 0; i < 5; i++ {
			in <- i
		}
	}()

	p := "@0\nMOV UP DOWN\n\n@1\nMOV UP DOWN\n"
	tis.Program(strings.NewReader(p))
	tis.Run()

	for i := 0; i < 5; i++ {
		x := <-out
		if x != i {
			t.Errorf("Expected %d, got %d", i, x)
		}
	}
}
