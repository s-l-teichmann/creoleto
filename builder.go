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

func (b *builder) up() {
	if b.current != nil {
		b.current = b.current.parent
	}
}

func (b *builder) down(n *node) {
	b.current = link(n, b.current)
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

	link(&node{nodeType: typ, value: content}, b.current)
}

func (b *builder) EnterParagraph(c *parser.ParagraphContext) {
	b.down(&node{nodeType: paragraphNode})
}

func (b *builder) ExitParagraph(c *parser.ParagraphContext) {
	b.up()
}

// text_first_element

func (b *builder) EnterText_firstelement(c *parser.Text_firstelementContext) {
	fmt.Fprintln(os.Stderr, "1. EnterText_firstelement")
}

func (b *builder) ExitText_firstelement(c *parser.Text_firstelementContext) {
	fmt.Fprintln(os.Stderr, "1. ExitText_firstelement")
}

// text_first_unformattedelement

func (b *builder) EnterText_first_unformattedelement(c *parser.Text_first_unformattedelementContext) {
	fmt.Fprintln(os.Stderr, "1. ExitText_first_unformattedelement")
}

func (b *builder) ExitText_first_unformattedelement(c *parser.Text_first_unformattedelementContext) {
	fmt.Fprintf(os.Stderr, "1. ExitText_firstunformattedelement: '%s'\n", c.GetText())
	link(text(c.GetText()), b.current)
}

// text_unformattedelement

func (b *builder) EnterText_unformattedelement(c *parser.Text_unformattedelementContext) {
	fmt.Fprintln(os.Stderr, "EnterText_unformattedelement")
}

func (b *builder) ExitText_unformattedelement(c *parser.Text_unformattedelementContext) {
	fmt.Fprintf(os.Stderr, "ExitText_unformattedelement '%s'\n", c.GetText())
	link(text(c.GetText()), b.current)
}

func (b *builder) EnterText_line(c *parser.Text_lineContext) {
	fmt.Fprintln(os.Stderr, "EnterText_line")
}

func (b *builder) ExitText_line(c *parser.Text_lineContext) {
	fmt.Fprintln(os.Stderr, "ExitText_line:", c.GetText())
	link(text("\n"), b.current)
}

func (b *builder) EnterText_formattedelement(c *parser.Text_formattedelementContext) {
	fmt.Fprintln(os.Stderr, "***EnterText_formattedelement")
	// type to be overwritten in ital_markup or bold_markup
	n := &node{nodeType: italicsNode}
	b.push(n)
	b.down(n)
}

func (b *builder) ExitText_formattedelement(c *parser.Text_formattedelementContext) {
	fmt.Fprintln(os.Stderr, "***ExitText_formattedelement")
	b.pop()
	b.up()
}

func (b *builder) EnterItal_markup(c *parser.Ital_markupContext) {
	fmt.Fprintln(os.Stderr, "//EnterItal_markup")
	if typeface, ok := b.top().(*node); ok {
		typeface.nodeType = italicsNode
	}
}

func (b *builder) EnterBold_markup(c *parser.Bold_markupContext) {
	fmt.Fprintln(os.Stderr, "**EnterBold_markup")
	if typeface, ok := b.top().(*node); ok {
		typeface.nodeType = boldNode
	}
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
