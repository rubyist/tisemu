package main

import (
	"testing"
	"time"
)

func TestInitValues(t *testing.T) {
	n := NewT21()

	if n.acc != 0 {
		t.Errorf("Expected ACC to be 0, got %d\n", n.acc)
	}

	if n.bak != 0 {
		t.Errorf("Expected BAK to be 0, got %d\n", n.bak)
	}
}

func TestSimpleMov(t *testing.T) {
	n := NewT21()
	n.Mov(4, ACC)

	if n.acc != 4 {
		t.Errorf("Expected ACC to be 4 after `MOV 4 ACC`, got %d\n", n.acc)
	}
}

func TestMovIn(t *testing.T) {
	n := NewT21()

	c := make(chan int, 1)
	n.up = c
	c <- 4

	n.Mov(UP, ACC)

	if n.acc != 4 {
		t.Errorf("Expected ACC to be 4 after `MOV UP ACC`, got %d\n", n.acc)
	}
}

func TestMovAccOut(t *testing.T) {
	n := NewT21()

	c := make(chan int, 1)
	n.down = c

	n.Mov(4, ACC)
	n.Mov(ACC, DOWN)

	v := <-c
	if v != 4 {
		t.Errorf("Expected to receive 4 from `MOV ACC DOWN`, got %d\n", v)
	}
}

func TestMovVal(t *testing.T) {
	n := NewT21()

	c := make(chan int, 1)
	n.down = c

	n.Mov(4, DOWN)

	v := <-c
	if v != 4 {
		t.Errorf("Expected to receive 4 from `MOV 4 DOWN`, got %d\n", v)
	}
}

func TestMovThrough(t *testing.T) {
	n := NewT21()

	up := make(chan int, 1)
	down := make(chan int, 1)

	n.up = up
	n.down = down

	up <- 4
	n.Mov(UP, DOWN)

	v := <-down
	if v != 4 {
		t.Errorf("Expected to receive 4 from `MOV UP DOWN`, got %d\n", v)
	}
}

func TestSwp(t *testing.T) {
	n := NewT21()
	n.Mov(4, ACC)
	n.Swp()

	if n.acc != 0 {
		t.Errorf("Expected ACC to be 0 after `SWP`, got %d\n", n.acc)
	}

	if n.bak != 4 {
		t.Errorf("Expected BAK to be 4 after `SWP`, got %d\n", n.bak)
	}

	n.Swp()

	if n.acc != 4 {
		t.Errorf("Expected ACC to be 4 after `SWP`, got %d\n", n.acc)
	}

	if n.bak != 0 {
		t.Errorf("Expected BAK to be 0 after `SWP`, got %d\n", n.bak)
	}
}

func TestSav(t *testing.T) {
	n := NewT21()
	n.Mov(4, ACC)
	n.Sav()

	if n.acc != 4 {
		t.Errorf("Expected ACC to remain 4 after `SAV`, got %d\n", n.acc)
	}

	if n.bak != 4 {
		t.Errorf("Expected BAK to be 4 after `SAV`, got %d\n", n.bak)
	}
}

func TestSimpleAdd(t *testing.T) {
	n := NewT21()
	n.Mov(4, ACC)
	n.Add(1)

	if n.acc != 5 {
		t.Errorf("Expected ACC to be 5 after `ADD 1`, got %d\n", n.acc)
	}
}

func TestSimpleSub(t *testing.T) {
	n := NewT21()
	n.Mov(4, ACC)
	n.Sub(1)

	if n.acc != 3 {
		t.Errorf("Expected ACC to be 3 after `SUB 1`, got %d\n", n.acc)
	}
}

func TestNeg(t *testing.T) {
	n := NewT21()
	n.Mov(4, ACC)
	n.Neg()

	if n.acc != -4 {
		t.Errorf("Expected ACC to be -4 after `NEG`, got %d\n", n.acc)
	}
}

func TestSimpleProgram(t *testing.T) {
	p := []Statement{
		{
			Op:  MOV,
			Src: 4,
			Dst: ACC,
		},
		{
			Op:  ADD,
			Src: 1,
		},
		{
			Op:  JRO,
			Src: 0,
		},
	}

	n := NewT21()
	n.Program(p)
	n.Run()

	time.Sleep(time.Millisecond * 20)

	if n.acc != 5 {
		t.Errorf("Expected ACC to be 5 after program run, got %d\n", n.acc)
	}
}
