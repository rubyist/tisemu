package main

type Node struct {
	Up    chan int
	Down  chan int
	Left  chan int
	Right chan int
	Acc   int
	Bak   int
	pc    int
	code  []command
}

func (n *Node) Run() {
	go func() {
		for {
			n.writeDown(n.readUp())
		}
	}()
}

func (n *Node) ConnectDown(neighbor *Node) {
	c := make(chan int)
	n.Down = c
	neighbor.Up = c
}

func (n *Node) readUp() int {
	return <-n.Up
}

func (n *Node) readDown() int {
	return <-n.Down
}

func (n *Node) readLeft() int {
	return <-n.Left
}

func (n *Node) readRight() int {
	return <-n.Right
}

func (n *Node) writeUp(v int) {
	n.Up <- v
}

func (n *Node) writeDown(v int) {
	n.Down <- v
}

func (n *Node) writeLeft(v int) {
	n.Left <- v
}

func (n *Node) writeRight(v int) {
	n.Right <- v
}

type command struct {
	Opcode int
	Value  int
}
