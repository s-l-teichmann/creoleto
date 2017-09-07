// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import "runtime"

type parser struct {
	lexer *lexer
}

func newParser() *parser {
	return &parser{}
}

func (p *parser) recover(err *error) {
	e := recover()
	if e == nil {
		return
	}
	if _, ok := e.(runtime.Error); ok {
		panic(e)
	}
	if p != nil {
		p.lexer.drain()
	}
	*err = e.(error)
}

func (p *parser) parse(input string) (doc *document, err error) {
	p.lexer = newLexer(input)
	defer p.recover(&err)
	return
}
