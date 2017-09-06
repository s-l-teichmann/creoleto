// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

type stateFn func(*lexer) stateFn

type lexer struct {
	input string
	state stateFn
}

func newLexer(input string) *lexer {
	return &lexer{
		input: input,
	}
}

func lexText(l *lexer) stateFn {
	return lexText
}
