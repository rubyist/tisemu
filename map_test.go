package main

import (
	"strings"
	"testing"
)

func TestMapWithDisplay(t *testing.T) {
	m := `4
3
T 18 30
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

	if nm.DisplayRows != 18 {
		t.Fatal("Expected display to be 18 rows, got %d", nm.DisplayRows)
	}

	if nm.DisplayCols != 30 {
		t.Fatal("Expected display to be 30 cols, got %d", nm.DisplayCols)
	}

	if nm.NodeMap[0][0] != T21Node {
		t.Fatalf("Expected T21 node, got %d", nm.NodeMap[0][0])
	}

	if nm.NodeMap[1][0] != T30Node {
		t.Fatalf("Expected T30 node, got %d", nm.NodeMap[1][0])
	}
}

func TestMapWithoutDisplay(t *testing.T) {
	m := `4
3
F
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

	if nm.Display {
		t.Fatal("Expected display to be disabled")
	}

	if nm.NodeMap[0][0] != T21Node {
		t.Fatalf("Expected T21 node, got %d", nm.NodeMap[0][0])
	}

	if nm.NodeMap[1][0] != T30Node {
		t.Fatalf("Expected T30 node, got %d", nm.NodeMap[1][0])
	}
}
