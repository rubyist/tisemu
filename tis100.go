package main

import (
	"io"
	"time"
)

type NodeType int
type NodeMap [][]NodeType

const (
	T21Node NodeType = iota
	T30Node
	T31Node
)

type Tis100 struct {
	clock <-chan time.Time
	nodes [][]*T21 // TODO: Need to support other node types
}

func NewTis100(nodes NodeMap) *Tis100 {
	// All rows need to be the same length
	l := len(nodes[0])
	for i := 1; i < len(nodes); i++ {
		if len(nodes[i]) != l {
			panic("all rows must be the same length")
		}
	}

	t := &Tis100{
		clock: time.Tick(time.Millisecond),
	}

	for _, list := range nodes {
		var cnodes []*T21

		for _, node := range list {
			switch node {
			case T21Node:
				cnodes = append(cnodes, NewT21())
			default:
				panic("uknown node type")
			}
		}

		t.nodes = append(t.nodes, cnodes)
	}

	// Connect nodes
	for r, row := range t.nodes {
		for i := 0; i < len(row)-1; i++ {
			row[i].ConnectRight(row[i+1])
		}

		for i := 0; i < len(row); i++ {
			if r < len(t.nodes)-1 {
				row[i].ConnectDown(t.nodes[r+1][i])
			}
		}
	}

	return t
}

func (t *Tis100) Input(in chan int, node int) {
	row, col := t.nodeOffsets(node)
	t.nodes[row][col].up = in
}

func (t *Tis100) Output(node int) chan int {
	row, col := t.nodeOffsets(node)

	out := make(chan int)
	t.nodes[row][col].down = out
	return out
}

func (t *Tis100) Program(r io.Reader) {
	p := NewParser(r)

	curNode := 0
	programs := make(map[int][]Statement, 0)

	for {
		stmt, err := p.Parse()
		if err != nil {
			panic(err)
		}
		if stmt.Op == EOF {
			break
		} else if stmt.Op == COMMENT {
			continue
		} else if stmt.Op == NODE {
			curNode = int(stmt.Src)
		} else {
			programs[curNode] = append(programs[curNode], stmt)
		}
	}

	for idx, program := range programs {
		row, col := t.nodeOffsets(idx)
		t.nodes[row][col].Program(program)
	}
}

func (t *Tis100) Run() {
	for _, nodes := range t.nodes {
		for _, node := range nodes {
			node.Run()
		}
	}

	go func() {
		for {
			<-t.clock
			for _, nodes := range t.nodes {
				for _, node := range nodes {
					node.tick()
				}
			}
		}
	}()
}

func (t *Tis100) nodeOffsets(node int) (int, int) {
	row := node / len(t.nodes[0])
	col := node % len(t.nodes[0])
	return row, col
}
