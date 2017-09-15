// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/s-l-teichmann/creoleto/parser"
)

type builder struct {
	*parser.BaseCreole10Listener
	current *node

	stack []interface{}
}

type headingBuilder struct {
	depth      int
	hasContent bool
}

func newBuilder() *builder {
	return &builder{
		current: &node{},
	}
}

func (b *builder) push(x interface{}) {
	b.stack = append(b.stack, x)
}

func (b *builder) pop() interface{} {
	l := len(b.stack)
	x := b.stack[l-1]
	b.stack[l-1] = nil
	b.stack = b.stack[:l-1]
	return x
}

func (b *builder) top() interface{} {
	return b.stack[len(b.stack)-1]
}

func (b *builder) EnterHeading(c *parser.HeadingContext) {
	b.push(&headingBuilder{})
}

func (b *builder) EnterHeading_markup(c *parser.Heading_markupContext) {
	if hb := b.top().(*headingBuilder); !hb.hasContent {
		hb.depth++
	}
}

func (b *builder) ExitHeading_content(c *parser.Heading_contentContext) {
	b.top().(*headingBuilder).hasContent = true
}

func (b *builder) ExitHeading(c *parser.HeadingContext) {

	hb := b.pop().(*headingBuilder)

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
		parent:   b.current,
		value:    content,
	}
	b.current.children = append(b.current.children, n)
}

func (b *builder) EnterParagraph(c *parser.ParagraphContext) {
	b.current = ndp(paragraphNode, b.current)
}

func (b *builder) ExitParagraph(c *parser.ParagraphContext) {
	if b.current != nil && b.current.nodeType == paragraphNode {
		b.current = b.current.parent
	}
}

func (b *builder) EnterText_unformattedelement(c *parser.Text_unformattedelementContext) {
	fmt.Fprintln(os.Stderr, "EnterText_unformattedelement")
}

func (b *builder) ExitText_unformattedelement(c *parser.Text_unformattedelementContext) {
	fmt.Fprintf(os.Stderr, "ExitText_unformattedelement '%s'\n", c.GetText())
}

func (b *builder) EnterText_line(c *parser.Text_lineContext) {
	fmt.Fprintln(os.Stderr, "EnterText_line")
}

func (b *builder) ExitText_line(c *parser.Text_lineContext) {
	fmt.Fprintln(os.Stderr, "ExitText_line:", c.GetText())
}

func (b *builder) EnterText_formattedelement(c *parser.Text_formattedelementContext) {
	fmt.Fprintln(os.Stderr, "***EnterText_formattedelement")
}

func (b *builder) ExitText_formattedelement(c *parser.Text_formattedelementContext) {
	fmt.Fprintln(os.Stderr, "***ExitText_formattedelement")
}

func (b *builder) EnterItal_markup(c *parser.Ital_markupContext) {
	fmt.Fprintln(os.Stderr, "//EnterItal_markup")
}

func (b *builder) parse(data string) (*document, error) {

	input := antlr.NewInputStream(data)
	lexer := parser.NewCreole10Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	pa := parser.NewCreole10Parser(stream)
	pa.AddParseListener(b)
	//pa.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//pa.BuildParseTrees = true
	pa.Start()
	//antlr.ParseTreeWalkerDefault.Walk(p, tree)

	// TODO: Write own ErrorListener.

	doc := &document{root: b.current}

	return doc, nil
}
