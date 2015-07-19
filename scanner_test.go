package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	var tests = []struct {
		s   string
		tok Token
		lit string
	}{
		// Specials
		{s: ``, tok: EOF},
		{s: `/`, tok: ILLEGAL, lit: `/`},
		{s: ` `, tok: WS, lit: " "},
		{s: "\t", tok: WS, lit: "\t"},
		{s: "\n", tok: WS, lit: "\n"},
		{s: "#foo", tok: COMMENT, lit: "#foo"},
		{s: "foo:", tok: LABEL, lit: "foo:"},
		{s: "@1", tok: NODE, lit: "@1"},

		{s: `42`, tok: NUMBER, lit: "42"},
		{s: `-42`, tok: NUMBER, lit: "-42"},
		{s: `+42`, tok: NUMBER, lit: "+42"},

		// Keywords
		{s: `MOV`, tok: MOV, lit: "MOV"},
		{s: `ADD`, tok: ADD, lit: "ADD"},
		{s: `NEG`, tok: NEG, lit: "NEG"},
		{s: `SUB`, tok: SUB, lit: "SUB"},
		{s: `SWP`, tok: SWP, lit: "SWP"},
		{s: `SAV`, tok: SAV, lit: "SAV"},
		{s: `JMP`, tok: JMP, lit: "JMP"},
		{s: `JGZ`, tok: JGZ, lit: "JGZ"},
		{s: `JEZ`, tok: JEZ, lit: "JEZ"},
		{s: `JLZ`, tok: JLZ, lit: "JLZ"},
		{s: `JNZ`, tok: JNZ, lit: "JNZ"},
		{s: `JRO`, tok: JRO, lit: "JRO"},
		{s: `UP`, tok: UP, lit: "UP"},
		{s: `DOWN`, tok: DOWN, lit: "DOWN"},
		{s: `LEFT`, tok: LEFT, lit: "LEFT"},
		{s: `RIGHT`, tok: RIGHT, lit: "RIGHT"},
		{s: `ANY`, tok: ANY, lit: "ANY"},
		{s: `LAST`, tok: LAST, lit: "LAST"},
		{s: `ACC`, tok: ACC, lit: "ACC"},
		{s: `NIL`, tok: NIL, lit: "NIL"},
		{s: `HCF`, tok: HCF, lit: "HCF"},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}

func TestScanMulti(t *testing.T) {
	type result struct {
		tok Token
		lit string
	}

	exp := []result{
		{tok: LABEL, lit: "L:"},
		{tok: MOV, lit: "MOV"},
		{tok: WS, lit: " "},
		{tok: NUMBER, lit: "42"},
		{tok: WS, lit: " "},
		{tok: DOWN, lit: "DOWN"},
		{tok: EOF, lit: ""},
	}

	v := `L:MOV 42 DOWN`
	s := NewScanner(strings.NewReader(v))

	var act []result
	for {
		tok, lit := s.Scan()
		act = append(act, result{tok, lit})
		if tok == EOF {
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
