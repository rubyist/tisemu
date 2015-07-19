package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var StandardMachine = &MachineMap{
	Cols:    4,
	Rows:    3,
	Display: false,
	NodeMap: [][]NodeType{
		{T21Node, T21Node, T21Node, T21Node},
		{T21Node, T21Node, T21Node, T21Node},
		{T21Node, T21Node, T21Node, T21Node},
	},
}

type MachineMap struct {
	Cols        int
	Rows        int
	Display     bool
	DisplayRows int
	DisplayCols int
	NodeMap     [][]NodeType
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
	enabled, dr, dc, err := readDisplay(buf)
	if err != nil {
		return nil, err
	}
	m.Display = enabled
	m.DisplayRows = dr
	m.DisplayCols = dc

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

func readDisplay(r *bufio.Reader) (bool, int, int, error) {
	l, err := r.ReadString('\n')
	if err != nil {
		return false, 0, 0, err
	}
	l = strings.TrimSpace(l)

	if l == "F" {
		return false, 0, 0, nil
	}

	parts := strings.Split(l, " ")
	if len(parts) != 3 {
		return false, 0, 0, fmt.Errorf("Invalid display description: %s", l)
	}

	rows, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, 0, 0, fmt.Errorf("Invalid row count %s", parts[1])
	}

	cols, err := strconv.Atoi(parts[2])
	if err != nil {
		return false, 0, 0, fmt.Errorf("Invalid col count %s", parts[2])
	}

	return true, rows, cols, nil
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
