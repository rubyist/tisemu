package main

import (
	"sync"
	"sync/atomic"
)

const STACKEMPTY = -1000

type T30 struct {
	up          chan int
	down        chan int
	left        chan int
	right       chan int
	stack       []int
	readTicker  chan interface{}
	writeTicker chan interface{}
	l           sync.Mutex
}

func NewT30() *T30 {
	return &T30{
		readTicker:  make(chan interface{}),
		writeTicker: make(chan interface{}),
	}
}

func (n *T30) push(v int) {
	n.stack = append(n.stack, v)
}

func (n *T30) peek() int {
	if len(n.stack) == 0 {
		return STACKEMPTY
	}
	return n.stack[len(n.stack)-1]
}

func (n *T30) pop() {
	if len(n.stack) == 0 {
		return
	}
	n.stack = n.stack[0 : len(n.stack)-1]
}

func (n *T30) Run() {
	// Service writers
	go func() {
		for {
			if atomic.LoadInt32(&hcf) == 1 {
				return
			}

			<-n.writeTicker
			n.l.Lock()
			var val int
			select {
			case val = <-n.up:
				n.push(val)
			case val = <-n.right:
				n.push(val)
			case val = <-n.down:
				n.push(val)
			case val = <-n.left:
				n.push(val)
			default:
			}
			n.l.Unlock()
		}
	}()

	// Service readers
	go func() {
		for {
			if atomic.LoadInt32(&hcf) == 1 {
				return
			}

			<-n.readTicker
			n.l.Lock()
			if val := n.peek(); val != STACKEMPTY {
				select {
				case n.up <- val:
					n.pop()
				case n.right <- val:
					n.pop()
				case n.down <- val:
					n.pop()
				case n.left <- val:
					n.pop()
				default: // no readers
				}
			}
			n.l.Unlock()
		}
	}()
}

func (n *T30) ConnectDown(neighbor *T30) {
	c := make(chan int)
	n.down = c
	neighbor.up = c
}

func (n *T30) ConnectRight(neighbor *T30) {
	c := make(chan int)
	n.right = c
	neighbor.left = c
}

func (n *T30) tick() {
	select {
	case n.readTicker <- 1:
	default:
	}

	select {
	case n.writeTicker <- 1:
	default:
	}
}

func (n *T30) Program(p []Statement) {
}
