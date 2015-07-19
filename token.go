package main

import "fmt"

type Token int

func (t Token) String() string {
	if t < 1000 {
		return fmt.Sprintf("%d", t)
	}
	return tokStrings[t]
}

const (
	ILLEGAL Token = iota + 1000
	EOF
	WS

	NODE
	COMMENT
	NOP
	LABEL
	IDENT
	NUMBER
	VAL

	MOV
	ADD
	SUB
	NEG
	SWP
	SAV

	JMP
	JGZ
	JEZ
	JLZ
	JNZ
	JRO

	UP
	DOWN
	LEFT
	RIGHT
	ACC
	NIL
	ANY
	LAST
	HCF
)

var keys = map[string]Token{
	"NOP":   NOP,
	"MOV":   MOV,
	"ADD":   ADD,
	"SUB":   SUB,
	"NEG":   NEG,
	"SWP":   SWP,
	"SAV":   SAV,
	"JMP":   JMP,
	"JGZ":   JGZ,
	"JEZ":   JEZ,
	"JLZ":   JLZ,
	"JNZ":   JNZ,
	"JRO":   JRO,
	"UP":    UP,
	"DOWN":  DOWN,
	"LEFT":  LEFT,
	"RIGHT": RIGHT,
	"ANY":   ANY,
	"LAST":  LAST,
	"ACC":   ACC,
	"NIL":   NIL,
	"HCF":   HCF,
}

var tokStrings = map[Token]string{
	NOP:   "NOP",
	MOV:   "MOV",
	ADD:   "ADD",
	SUB:   "SUB",
	NEG:   "NEG",
	SWP:   "SWP",
	SAV:   "SAV",
	JMP:   "JMP",
	JGZ:   "JGZ",
	JEZ:   "JEZ",
	JLZ:   "JLZ",
	JNZ:   "JNZ",
	JRO:   "JRO",
	UP:    "UP",
	DOWN:  "DOWN",
	LEFT:  "LEFT",
	RIGHT: "RIGHT",
	ANY:   "ANY",
	LAST:  "LAST",
	ACC:   "ACC",
	NIL:   "NIL",
	HCF:   "HCF",
}

var eof = rune(0)
