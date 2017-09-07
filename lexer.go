// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"fmt"
	"unicode/utf8"
)

type itemType int

type item struct {
	typ  itemType
	pos  int
	val  string
	line int
}

const (
	itemError itemType = iota
	itemEOF
)

const eof = -1

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	case len(i.val) > 10:
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf(i.val)
}

type stateFn func(*lexer) stateFn

type lexer struct {
	input   string
	state   stateFn
	pos     int
	start   int
	width   int
	lastPos int
	line    int

	items chan item
}

func newLexer(input string) *lexer {
	l := &lexer{
		input: input,
		line:  1,
		items: make(chan item),
	}
	go l.run()
	return l
}

func (l *lexer) drain() {
	for range l.items {
	}
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += w
	if r == '\n' {
		l.line++
	}
	return r
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

func (l *lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos], l.line}
	// TODO: Count newlines for text containing items.
	l.start = l.pos
}

func (l *lexer) nextItem() item {
	item := <-l.items
	l.lastPos = item.pos
	return item
}

func lexText(l *lexer) stateFn {
	return lexText
}
