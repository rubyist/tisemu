package main

import (
	"testing"
	"time"
)

func TestInitValues(t *testing.T) {
	n := &T21{}

	if n.acc != 0 {
		t.Errorf("Expected ACC to be 0, got %d\n", n.acc)
	}

	if n.bak != 0 {
		t.Errorf("Expected BAK to be 0, got %d\n", n.bak)
	}
}

func TestSimpleMov(t *testing.T) {
	n := &T21{}
	n.Mov(4, ACC)

	if n.acc != 4 {
		t.Errorf("Expected ACC to be 4 after `MOV 4 ACC`, got %d\n", n.acc)
	}
}

func TestSwp(t *testing.T) {
	n := &T21{}
	n.Mov(4, ACC)
	n.Swp()

	if n.acc != 0 {
		t.Errorf("Expected ACC to be 0 after `SWP`, got %d\n", n.acc)
	}

	if n.bak != 4 {
		t.Errorf("Expected BAK to be 4 after `SWP`, got %d\n", n.bak)
	}
}

func TestSav(t *testing.T) {
	n := &T21{}
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
	n := &T21{}
	n.Mov(4, ACC)
	n.Add(1)

	if n.acc != 5 {
		t.Errorf("Expected ACC to be 5 after `ADD 1`, got %d\n", n.acc)
	}
}

func TestSimpleSub(t *testing.T) {
	n := &T21{}
	n.Mov(4, ACC)
	n.Sub(1)

	if n.acc != 3 {
		t.Errorf("Expected ACC to be 3 after `SUB 1`, got %d\n", n.acc)
	}
}

func TestNeg(t *testing.T) {
	n := &T21{}
	n.Mov(4, ACC)
	n.Neg()

	if n.acc != -4 {
		t.Errorf("Expected ACC to be -4 after `NEG`, got %d\n", n.acc)
	}
}

func TestSimpleProgram(t *testing.T) {
	var p program
	p[0] = instruction{Op: MOV, Src: 4, Dst: ACC}
	p[1] = instruction{Op: ADD, Src: 1}
	p[2] = instruction{Op: JRO, Src: 0}

	n := &T21{}
	n.p = p
	n.Run()

	time.Sleep(time.Millisecond * 20)

	if n.acc != 5 {
		t.Errorf("Expected ACC to be 5 after program run, got %d\n", n.acc)
	}
}
