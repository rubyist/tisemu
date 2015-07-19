package main

import (
	"testing"
	"time"
)

func TestT30(t *testing.T) {
	n := NewT30()

	in := make(chan int)
	n.left = in

	out := make(chan int)
	n.down = out

	n.Run()

	term := make(chan int)

	go func() {
		for {
			select {
			case <-term:
				return
			default:
				time.Sleep(time.Millisecond * 5)
				n.tick()
			}
		}
	}()

	n.left <- 23
	n.left <- 42

	v := <-n.down
	if v != 42 {
		t.Fatalf("expected to read 42, got %d", v)
	}

	v = <-n.down
	if v != 23 {
		t.Fatalf("expected to read 23, got %d", v)
	}

	close(term)
}
