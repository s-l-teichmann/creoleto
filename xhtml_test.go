// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"bytes"
	"testing"
)

func TestStandalone(t *testing.T) {
	doc := document{
		root: &node{
			nodeType: heading1Node,
			value:    "Hello",
		},
	}

	var buf bytes.Buffer

	if err := exportXHTML(&doc, &buf, true); err != nil {
		t.Fatalf("error: %v\n", err)
	}

	const want = xhtmlHeader + "<h1>Hello</h1>" + xhtmlFooter

	if got := buf.String(); got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestElements(t *testing.T) {

	cases := []testCase{{
		have: text("Hello <&>"),
		want: "Hello &lt;&amp;&gt;",
	}, {
		have: &node{nodeType: horizontalLineNode},
		want: "<hr/>",
	}, {
		have: &node{nodeType: lineBreakNode},
		want: "<br/>",
	}, {
		have: nd(italicsNode, text("Hello")),
		want: "<i>Hello</i>",
	}, {
		have: nd(boldNode, text("Hello")),
		want: "<strong>Hello</strong>",
	}, {
		have: nd(strikeNode, text("Hello")),
		want: "<del>Hello</del>",
	}, {
		have: nd(underlinedNode, text("Hello")),
		want: "<em>Hello</em>",
	}, {
		have: nd(subscriptNode, text("Hello")),
		want: "<sub>Hello</sub>",
	}, {
		have: nd(superscriptNode, text("Hello")),
		want: "<sup>Hello</sup>",
	}, {
		have: nd(noWikiNode, text("Hello")),
		want: "<pre><tt>Hello</tt></pre>",
	}, {
		have: nd(noWikiInlineNode, text("Hello")),
		want: "<tt>Hello</tt>",
	}, {
		have: &node{nodeType: heading1Node},
		want: "<h1></h1>",
	}, {
		have: &node{nodeType: heading2Node},
		want: "<h2></h2>",
	}, {
		have: &node{nodeType: heading3Node},
		want: "<h3></h3>",
	}, {
		have: &node{nodeType: heading4Node},
		want: "<h4></h4>",
	}, {
		have: &node{nodeType: heading5Node},
		want: "<h5></h5>",
	}, {
		have: &node{nodeType: heading6Node},
		want: "<h6></h6>",
	}, {
		have: &node{nodeType: paragraphNode},
		want: "<p></p>",
	}, {
		have: &node{
			nodeType: linkNode,
			value:    "http://example.com",
			children: []*node{text("Hello")},
		},
		want: `<a href="http://example.com">Hello</a>`,
	}, {
		have: &node{
			nodeType: imageNode,
			value: &image{
				src: "http://example.com/image.png",
			},
		},
		want: `<img src="http://example.com/image.png"></img>`,
	}, {
		have: &node{
			nodeType: imageNode,
			value: &image{
				src: "http://example.com/image.png",
				alt: "An image",
			},
		},
		want: `<img src="http://example.com/image.png" alt="An image"></img>`,
	}, {
		have: nd(unorderedListNode,
			nd(listItemNode, text("Hello")),
			nd(listItemNode, text("World"))),
		want: `<ul><li>Hello</li><li>World</li></ul>`,
	}, {
		have: nd(orderedListNode,
			nd(listItemNode, text("Hello")),
			nd(listItemNode, text("World"))),
		want: `<ol><li>Hello</li><li>World</li></ol>`,
	}, {
		have: nd(tableNode,
			nd(tableHeaderRowNode,
				nd(tableHeaderCellNode, text("header1")),
				nd(tableHeaderCellNode, text("header2"))),
			nd(tableRowNode,
				nd(tableCellNode, text("Hello")),
				nd(tableCellNode, text("World")))),
		want: "<table>" +
			"<tr>" +
			"<th>header1</th>" +
			"<th>header2</th>" +
			"</tr>" +
			"<tr>" +
			"<td>Hello</td>" +
			"<td>World</td>" +
			"</tr>" +
			"</table>",
	}}

	runCases(cases, exportXHTML, t)
}
