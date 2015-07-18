package main

type T21 struct {
	up     chan int
	down   chan int
	left   chan int
	right  chan int
	acc    int
	bak    int
	pc     int
	p      []Statement
	term   chan interface{}
	ticker chan interface{}
}

func NewT21() *T21 {
	t21 := &T21{
		ticker: make(chan interface{}),
	}
	return t21
}

func (n *T21) Program(p []Statement) {
	n.p = p
}

func (n *T21) Nop() {
}

func (n *T21) Mov(src, dst Token) {
	val := int(src)
	switch src {
	case UP:
		val = n.readUp()
	case DOWN:
		val = n.readDown()
	case LEFT:
		val = n.readLeft()
	case RIGHT:
		val = n.readRight()
	case ANY:
		val = n.readAny()
	case ACC:
		val = n.acc
	}

	switch dst {
	case ACC:
		n.acc = val
	case UP:
		n.writeUp(val)
	case DOWN:
		n.writeDown(val)
	case LEFT:
		n.writeLeft(val)
	case RIGHT:
		n.writeRight(val)
	case ANY:
		n.writeAny(val)
	default:
		panic("unknown destination")
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

func (n *T21) Add(src Token) {
	val := int(src)

	switch src {
	case ACC:
		val = n.acc
	case LEFT:
		val = n.readLeft()
	case RIGHT:
		val = n.readRight()
	case UP:
		val = n.readUp()
	case DOWN:
		val = n.readDown()
	}

	n.acc += val
}

func (n *T21) Sub(src Token) {
	val := int(src)

	switch src {
	case ACC:
		val = n.acc
	case LEFT:
		val = n.readLeft()
	case RIGHT:
		val = n.readRight()
	case UP:
		val = n.readUp()
	case DOWN:
		val = n.readDown()
	}

	n.acc -= val
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

func (n *T21) Jro(src Token) {
	if int(src) == 0 {
		close(n.term)
	}
}

func (n *T21) Run() {
	if len(n.p) == 0 {
		n.p = []Statement{
			{
				Op: NOP,
			},
		}
	}

	n.term = make(chan interface{})

	go func() {
		for {
			<-n.ticker
			select {
			case <-n.term:
				return
			default:
			}

			if n.pc > len(n.p)-1 { // TODO Should limit to 16
				n.pc = 0
			}

			command := n.p[n.pc]
			switch command.Op {
			case NOP:
			case MOV:
				n.Mov(command.Src, command.Dst)
			case SWP:
				n.Swp()
			case SAV:
				n.Sav()
			case ADD:
				n.Add(command.Src)
			case SUB:
				n.Sub(command.Src)
			case NEG:
				n.Neg()
			case JMP:
				n.Jmp(command.Label)
			case JEZ:
				n.Jez(command.Label)
			case JNZ:
				n.Jnz(command.Label)
			case JGZ:
				n.Jgz(command.Label)
			case JLZ:
				n.Jlz(command.Label)
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

func (n *T21) ConnectLeft(neighbor *T21) {
	c := make(chan int)
	n.left = c
	neighbor.right = c
}

func (n *T21) ConnectRight(neighbor *T21) {
	c := make(chan int)
	n.right = c
	neighbor.left = c
}

func (n *T21) tick() {
	select {
	case n.ticker <- 1:
	default:
	}
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

func (n *T21) readAny() int {
	select {
	case v := <-n.up:
		return v
	case v := <-n.down:
		return v
	case v := <-n.left:
		return v
	case v := <-n.right:
		return v
	}
}

func (n *T21) writeAny(v int) {
	select {
	case n.up <- v:
	case n.down <- v:
	case n.left <- v:
	case n.right <- v:
	}
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
