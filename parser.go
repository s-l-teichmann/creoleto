// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	prs "github.com/s-l-teichmann/creoleto/parser"
)

type parser struct {
	*prs.BaseCreole10Listener
}

func newParser() *parser {
	return &parser{}
}

func (p *parser) parse(data string) (*document, error) {

	input := antlr.NewInputStream(data)
	lexer := prs.NewCreole10Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	pa := prs.NewCreole10Parser(stream)
	pa.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	pa.BuildParseTrees = true

	//antlr.ParseTreeWalkerDefault.Walk(p, nil)

	return nil, nil
}
