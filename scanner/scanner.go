package scanner

import (
	"github.com/AlanValdevenito/monkey-interpreter/token"
)

type Scanner struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Scanner {
	s := &Scanner{input: input}
	s.readChar()
	return s
}
