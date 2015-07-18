package main

import (
	"fmt"
	"io"
	"strconv"
)

type Statement struct {
	Op    Token
	Src   Token
	Dst   Token
	Label string
}

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (Statement, error) {
	var stmt Statement

	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case EOF:
		return Statement{Op: EOF}, nil
	case NODE:
		n, _ := strconv.Atoi(lit[1:len(lit)])
		return Statement{Op: NODE, Src: Token(n)}, nil
	case LABEL:
		return Statement{Op: LABEL, Label: lit}, nil
	case NOP:
		return Statement{Op: NOP}, nil
	case MOV:
		return p.parseMovStatement()
	case ADD, SUB:
		return p.parseMathStatement(tok)
	case NEG:
		return Statement{Op: NEG}, nil
	case JMP, JEZ, JLZ, JGZ, JNZ:
		return p.parseLabeledJump(tok)
	case JRO:
		return p.parseJro()
	}

	return stmt, fmt.Errorf("bad statement: %s", lit)
}

func (p *Parser) parseMovStatement() (Statement, error) {
	stmt := Statement{Op: MOV}

	tok, lit := p.scanIgnoreWhitespace()
	if !isValidSrc(tok) {
		return stmt, newParseError(lit)
	}

	if tok == NUMBER {
		v, _ := strconv.Atoi(lit)
		stmt.Src = Token(v)
	} else {
		stmt.Src = tok
	}

	tok, lit = p.scanIgnoreWhitespace()
	if !isValidDst(tok) {
		return stmt, newParseError(lit)
	}
	stmt.Dst = tok

	return stmt, nil
}

func (p *Parser) parseMathStatement(op Token) (Statement, error) {
	stmt := Statement{Op: op}

	tok, lit := p.scanIgnoreWhitespace()
	if !isValidSrc(tok) {
		return stmt, newParseError(lit)
	}

	if tok == NUMBER {
		v, _ := strconv.Atoi(lit)
		stmt.Src = Token(v)
	} else {
		stmt.Src = tok
	}

	return stmt, nil
}

func (p *Parser) parseLabeledJump(op Token) (Statement, error) {
	stmt := Statement{Op: op}

	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return stmt, newParseError(lit)
	}

	stmt.Label = lit
	return stmt, nil
}

func (p *Parser) parseJro() (Statement, error) {
	stmt := Statement{Op: JRO}

	tok, lit := p.scanIgnoreWhitespace()
	if tok != NUMBER {
		return stmt, newParseError(lit)
	}

	v, _ := strconv.Atoi(lit)
	stmt.Src = Token(v)

	return stmt, nil
}

func (p *Parser) scanIgnoreWhitespace() (Token, string) {
	tok, lit := p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return tok, lit
}

func (p *Parser) scan() (Token, string) {
	return p.s.Scan()
}

func isValidSrc(tok Token) bool {
	if tok >= -999 && tok <= 999 {
		return true
	}

	switch tok {
	case NUMBER, ACC, UP, DOWN, LEFT, RIGHT, ANY, LAST:
		return true
	}
	return false
}

func isValidDst(tok Token) bool {
	// UP, DOWN, LEFT, RIGHT, NIL
	switch tok {
	case ACC, UP, DOWN, LEFT, RIGHT, ANY, LAST, NIL:
		return true
	}
	return false
}

func newParseError(found string) error {
	return fmt.Errorf("error: %s", found)
}
