package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type MachineMap struct {
	Cols    int
	Rows    int
	Display bool
	NodeMap [][]NodeType
}

func NewMachineMap(r io.Reader) (*MachineMap, error) {
	buf := bufio.NewReader(r)

	m := &MachineMap{}

	// Number of columns
	cols, err := readInt(buf)
	if err != nil {
		return nil, err
	}
	m.Cols = cols

	// Number of Rows
	rows, err := readInt(buf)
	if err != nil {
		return nil, err
	}
	m.Rows = rows

	// Has a display?
	d, err := readBool(buf)
	if err != nil {
		return nil, err
	}
	m.Display = d

	// Node types
	for i := 0; i < rows; i++ {
		var nodes []NodeType
		for j := 0; j < cols; j++ {
			t, err := readNodeType(buf)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, t)
		}
		m.NodeMap = append(m.NodeMap, nodes)
	}

	return m, nil
}

func readInt(r *bufio.Reader) (int, error) {
	l, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(strings.TrimSpace(l))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func readBool(r *bufio.Reader) (bool, error) {
	l, err := r.ReadString('\n')
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(l) == "T" {
		return true, nil
	}
	return false, nil
}

func readNodeType(r *bufio.Reader) (NodeType, error) {
	l, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	switch strings.TrimSpace(l) {
	case "T21":
		return T21Node, nil
	case "T30":
		return T30Node, nil
	default:
		return 0, fmt.Errorf("Uknown node type: %s", l)
	}
}