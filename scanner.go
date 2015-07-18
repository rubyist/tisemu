package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (Token, string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}

	if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '@':
		s.unread()
		return s.scanNode()
	case '-', '+':
		s.unread()
		return s.scanNumber()
	case '#':
		s.unread()
		return s.scanComment()

	}

	return ILLEGAL, string(ch)

}

func (s *Scanner) scanWhitespace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanIdent() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != ':' {
			s.unread()
			break
		} else if ch == ':' {
			buf.WriteRune(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	str := buf.String()

	if str[len(str)-1] == ':' {
		return LABEL, str[0:len(str)]
	}

	if op, ok := keys[strings.ToUpper(str)]; ok {
		return op, str
	}

	return IDENT, str
}

func (s *Scanner) scanNumber() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return NUMBER, buf.String()
}

func (s *Scanner) scanNode() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return NODE, buf.String()
}

func (s *Scanner) scanComment() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '\n' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return COMMENT, buf.String()
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}

func isWhitespace(ch rune) bool {
	if ch == ' ' || ch == '\t' || ch == '\n' {
		return true
	}
	return false
}

func isLetter(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

func isDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
