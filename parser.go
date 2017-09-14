// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	prs "github.com/s-l-teichmann/creoleto/parser"
)

type parser struct {
	*prs.BaseCreole10Listener
	current *node

	stack []interface{}
}

type headingBuilder struct {
	depth      int
	hasContent bool
}

func newParser() *parser {
	return &parser{
		current: &node{},
	}
}

func (p *parser) push(x interface{}) {
	p.stack = append(p.stack, x)
}

func (p *parser) pop() interface{} {
	l := len(p.stack)
	x := p.stack[l-1]
	p.stack[l-1] = nil
	p.stack = p.stack[:l-1]
	return x
}

func (p *parser) top() interface{} {
	return p.stack[len(p.stack)-1]
}

func (p *parser) EnterHeading(c *prs.HeadingContext) {
	p.push(&headingBuilder{})
}

func (p *parser) EnterHeading_markup(c *prs.Heading_markupContext) {
	if hb := p.top().(*headingBuilder); !hb.hasContent {
		hb.depth++
	}
}

func (p *parser) ExitHeading_content(c *prs.Heading_contentContext) {
	p.top().(*headingBuilder).hasContent = true
}

func (p *parser) ExitHeading(c *prs.HeadingContext) {

	hb := p.pop().(*headingBuilder)

	trim := strings.Repeat("=", hb.depth)
	content := strings.TrimSpace(c.GetText())
	content = strings.TrimSuffix(strings.TrimPrefix(content, trim), trim)
	content = strings.TrimSpace(content)

	var typ nodeType
	switch hb.depth {
	case 1:
		typ = heading1Node
	case 2:
		typ = heading2Node
	case 3:
		typ = heading3Node
	case 4:
		typ = heading4Node
	case 5:
		typ = heading5Node
	default:
		typ = heading6Node
	}

	n := &node{
		nodeType: typ,
		parent:   p.current,
		value:    content,
	}
	p.current.children = append(p.current.children, n)
}

func (p *parser) EnterParagraph(c *prs.ParagraphContext) {
	p.current = ndp(paragraphNode, p.current)
}

func (p *parser) ExitParagraph(c *prs.ParagraphContext) {
	if p.current != nil && p.current.nodeType == paragraphNode && p.current.parent != nil {
		p.current = p.current.parent
	}
}

func (p *parser) parse(data string) (*document, error) {

	input := antlr.NewInputStream(data)
	lexer := prs.NewCreole10Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	pa := prs.NewCreole10Parser(stream)
	pa.AddParseListener(p)
	pa.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//pa.BuildParseTrees = true
	pa.Start()
	//antlr.ParseTreeWalkerDefault.Walk(p, tree)

	// TODO: Write own ErrorListener.

	doc := &document{root: p.current}

	return doc, nil
}
