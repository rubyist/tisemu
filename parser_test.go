package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParserParse(t *testing.T) {
	var tests = []struct {
		s    string
		stmt Statement
		err  string
	}{
		{
			s: ``,
			stmt: Statement{
				Op: EOF,
			},
		},
		{
			s: `@2`,
			stmt: Statement{
				Op:  NODE,
				Src: 2,
			},
		},
		{
			s: `NOP`,
			stmt: Statement{
				Op: NOP,
			},
		},
		{
			s: `MOV 23 DOWN`,
			stmt: Statement{
				Op:  MOV,
				Src: 23,
				Dst: DOWN,
			},
		},
		{
			s: `MOV UP DOWN`,
			stmt: Statement{
				Op:  MOV,
				Src: UP,
				Dst: DOWN,
			},
		},
		{
			s: `MOV ACC DOWN`,
			stmt: Statement{
				Op:  MOV,
				Src: ACC,
				Dst: DOWN,
			},
		},
		{
			s: `MOV DOWN ACC`,
			stmt: Statement{
				Op:  MOV,
				Src: DOWN,
				Dst: ACC,
			},
		},
		{
			s: `MOV DOWN NIL`,
			stmt: Statement{
				Op:  MOV,
				Src: DOWN,
				Dst: NIL,
			},
		},
		{
			s: `ADD UP`,
			stmt: Statement{
				Op:  ADD,
				Src: UP,
			},
		},
		{
			s: `ADD 23`,
			stmt: Statement{
				Op:  ADD,
				Src: 23,
			},
		},
		{
			s: `ADD -23`,
			stmt: Statement{
				Op:  ADD,
				Src: -23,
			},
		},
		{
			s: `SUB UP`,
			stmt: Statement{
				Op:  SUB,
				Src: UP,
			},
		},
		{
			s: `SUB 23`,
			stmt: Statement{
				Op:  SUB,
				Src: 23,
			},
		},
		{
			s: `SUB -23`,
			stmt: Statement{
				Op:  SUB,
				Src: -23,
			},
		},
		{
			s: `NEG`,
			stmt: Statement{
				Op: NEG,
			},
		},
		{
			s: `FOO:`,
			stmt: Statement{
				Op:    LABEL,
				Label: "FOO:",
			},
		},
		{
			s: `JMP FOO`,
			stmt: Statement{
				Op:    JMP,
				Label: "FOO",
			},
		},
		{
			s: `JEZ FOO`,
			stmt: Statement{
				Op:    JEZ,
				Label: "FOO",
			},
		},
		{
			s: `JLZ FOO`,
			stmt: Statement{
				Op:    JLZ,
				Label: "FOO",
			},
		},
		{
			s: `JGZ FOO`,
			stmt: Statement{
				Op:    JGZ,
				Label: "FOO",
			},
		},
		{
			s: `JNZ FOO`,
			stmt: Statement{
				Op:    JNZ,
				Label: "FOO",
			},
		},
		{
			s: `JRO 1`,
			stmt: Statement{
				Op:  JRO,
				Src: 1,
			},
		},
	}

	for i, tt := range tests {
		stmt, err := NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n exp=%s\n got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

func TestParseProgram(t *testing.T) {
	prog := "@0\n\n@1\nMOV UP RIGHT\n@2\nMOV LEFT ACC\nSUB UP\nL: MOV ACC DOWN\nJMP L\n"

	p := NewParser(strings.NewReader(prog))

	exp := []Statement{
		{
			Op:  NODE,
			Src: 0,
		},
		{
			Op:  NODE,
			Src: 1,
		},
		{
			Op:  MOV,
			Src: UP,
			Dst: RIGHT,
		},
		{
			Op:  NODE,
			Src: 2,
		},
		{
			Op:  MOV,
			Src: LEFT,
			Dst: ACC,
		},
		{
			Op:  SUB,
			Src: UP,
		},
		{
			Op:    LABEL,
			Label: "L:",
		},
		{
			Op:  MOV,
			Src: ACC,
			Dst: DOWN,
		},
		{
			Op:    JMP,
			Label: "L",
		},
		{
			Op: EOF,
		},
	}

	var act []Statement
	for {
		stmt, err := p.Parse()
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		act = append(act, stmt)
		if stmt.Op == EOF {
			break
		}
	}

	if len(exp) != len(act) {
		t.Fatalf("token count mismatch: exp=%d, got=%d", len(exp), len(act))
	}

	for i := range exp {
		if !reflect.DeepEqual(exp[i], act[i]) {
			t.Fatalf("%d. token mismatch:\n\nexp=%#v\n\ngot=%#v", i, exp[i], act[i])
		}
	}
}

func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
