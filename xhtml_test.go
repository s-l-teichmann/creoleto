// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"bytes"
	"encoding/xml"
	"testing"
)

// linkChildren repairs the parent/child relations.
func linkChildren(n, p *node) {
	if n != nil {
		n.parent = p
		for _, c := range n.children {
			linkChildren(c, n)
		}
	}
}

func text(txt string) *node {
	return &node{
		nodeType: textNode,
		value:    txt,
	}
}

func nd(typ nodeType, children []*node) *node {
	return &node{
		nodeType: typ,
		children: children,
	}
}

func TestStandalone(t *testing.T) {
	doc := document{
		root: nd(heading1Node, []*node{text("Hello")}),
	}

	var buf bytes.Buffer

	if err := exportXHTML(&doc, &buf, true); err != nil {
		t.Fatalf("error: %v\n", err)
	}

	const want = xml.Header +
		"<html>\n<body>\n" +
		"<h1>Hello</h1>\n" +
		"</body>\n</html>\n"

	if got := buf.String(); got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestElements(t *testing.T) {

	cases := [...]struct {
		have *node
		want string
	}{{
		have: text("Hello <&>"),
		want: "Hello &lt;&amp;&gt;",
	}, {
		have: &node{nodeType: horizontalLineNode},
		want: "<hr/>",
	}, {
		have: &node{nodeType: lineBreakNode},
		want: "<br/>",
	}, {
		have: nd(italicsNode, []*node{text("Hello")}),
		want: "<i>Hello</i>",
	}, {
		have: nd(boldNode, []*node{text("Hello")}),
		want: "<strong>Hello</strong>",
	}, {
		have: nd(strikeNode, []*node{text("Hello")}),
		want: "<del>Hello</del>",
	}, {
		have: nd(underlinedNode, []*node{text("Hello")}),
		want: "<em>Hello</em>",
	}, {
		have: nd(subscriptNode, []*node{text("Hello")}),
		want: "<sub>Hello</sub>",
	}, {
		have: nd(superscriptNode, []*node{text("Hello")}),
		want: "<sup>Hello</sup>",
	}, {
		have: nd(noWikiNode, []*node{text("Hello")}),
		want: "<pre><tt>Hello</tt></pre>",
	}, {
		have: nd(noWikiInlineNode, []*node{text("Hello")}),
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
		want: `<img src="http://example.com/image.png" src="An image"></img>`,
	}, {
		have: &node{
			nodeType: unorderedListNode,
			children: []*node{
				nd(listItemNode, []*node{text("Hello")}),
				nd(listItemNode, []*node{text("World")}),
			},
		},
		want: `<ul><li>Hello</li><li>World</li></ul>`,
	}, {
		have: &node{
			nodeType: orderedListNode,
			children: []*node{
				nd(listItemNode, []*node{text("Hello")}),
				nd(listItemNode, []*node{text("World")}),
			},
		},
		want: `<ol><li>Hello</li><li>World</li></ol>`,
	}, {
		have: &node{
			nodeType: tableNode,
			children: []*node{
				&node{
					nodeType: tableHeaderRowNode,
					children: []*node{
						nd(tableHeaderCellNode, []*node{text("header1")}),
						nd(tableHeaderCellNode, []*node{text("header2")}),
					},
				},
				&node{
					nodeType: tableRowNode,
					children: []*node{
						nd(tableCellNode, []*node{text("Hello")}),
						nd(tableCellNode, []*node{text("World")}),
					},
				},
			},
		},
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

	var buf bytes.Buffer
	var doc document

	for i := range cases {
		c := &cases[i]
		buf.Reset()
		linkChildren(c.have, nil)
		doc.root = c.have
		if err := exportXHTML(&doc, &buf, false); err != nil {
			t.Fatalf("failed: %v", err)
		}

		if got := buf.String(); got != c.want {
			t.Errorf("got '%s': want '%s'\n", got, c.want)
		}
	}
}
