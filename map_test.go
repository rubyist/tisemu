package main

import (
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	m := `4
3
T
T21
T21
T21
T21
T30
T21
T21
T21
T21
T21
T21
T21
`

	nm, err := NewMachineMap(strings.NewReader(m))
	if err != nil {
		t.Fatalf("Error creating machine map: %s", err)
	}

	if nm.Cols != 4 {
		t.Fatalf("Expected 4 columns, got %d", nm.Cols)
	}

	if nm.Rows != 3 {
		t.Fatalf("Expected 3 rows, got %d", nm.Rows)
	}

	if !nm.Display {
		t.Fatal("Expected display to be enabled")
	}

	if nm.NodeMap[0][0] != T21Node {
		t.Fatalf("Expected T21 node, got %d", nm.NodeMap[0][0])
	}

	if nm.NodeMap[1][0] != T30Node {
		t.Fatalf("Expected T30 node, got %d", nm.NodeMap[1][0])
	}
}
