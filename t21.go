package main

type T21 struct {
	up    chan int
	down  chan int
	left  chan int
	right chan int
	acc   int
	bak   int
	pc    int
	code  []command
}

type Port int

const (
	UP = iota + 1000
	DOWN
	LEFT
	RIGHT
	ACC
	NIL
)

func (n *T21) Nop() {
}

func (n *T21) Mov(src, dst Port) {
	switch dst {
	case ACC:
		n.acc = int(src)
	}
}

func (n *T21) Swp() {
	tmp := n.bak
	n.bak = n.acc
	n.acc = tmp
}

func (n *T21) Sav() {
	n.bak = n.acc
}

func (n *T21) Add(src Port) {
	n.acc += int(src)
}

func (n *T21) Sub(src Port) {
	n.acc -= int(src)
}

func (n *T21) Neg() {
	n.acc *= -1
}

func (n *T21) Jmp(label string) {
}

func (n *T21) Jez(label string) {
}

func (n *T21) Jnz(label string) {
}

func (n *T21) Jgz(label string) {
}

func (n *T21) Jlz(label string) {
}

func (n *T21) Jro(src int) {
}

func (n *T21) Run() {
	go func() {
		for {
			n.writeDown(n.readUp())
		}
	}()
}

func (n *T21) ConnectDown(neighbor *T21) {
	c := make(chan int)
	n.down = c
	neighbor.up = c
}

func (n *T21) readUp() int {
	return <-n.up
}

func (n *T21) readDown() int {
	return <-n.down
}

func (n *T21) readLeft() int {
	return <-n.left
}

func (n *T21) readRight() int {
	return <-n.right
}

func (n *T21) writeUp(v int) {
	n.up <- v
}

func (n *T21) writeDown(v int) {
	n.down <- v
}

func (n *T21) writeLeft(v int) {
	n.left <- v
}

func (n *T21) writeRight(v int) {
	n.right <- v
}

type command struct {
	Opcode int
	Value  int
}
