package main

type T21 struct {
	up    chan int
	down  chan int
	left  chan int
	right chan int
	acc   int
	bak   int
	pc    int
	p     program
	term  chan interface{}
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

const (
	NOP = iota
	MOV
	SWP
	SAV
	ADD
	SUB
	NEG
	JMP
	JEZ
	JNZ
	JGZ
	JLZ
	JRO
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

func (n *T21) Jro(src Port) {
	if int(src) == 0 {
		close(n.term)
	}
}

func (n *T21) Run() {
	n.term = make(chan interface{})

	go func() {
		for {
			select {
			case <-n.term:
				return
			default:
			}

			if n.pc > 15 {
				n.pc = 0
			}

			command := n.p[n.pc]
			switch command.Op {
			case NOP:
			case MOV:
				n.Mov(command.Src, command.Dst)
			case ADD:
				n.Add(command.Src)
			case JRO:
				n.Jro(command.Src)
			default:
			}

			n.pc++
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

type instruction struct {
	Op    int
	Src   Port
	Dst   Port
	Label string
}

type program [16]instruction
