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

func TestStandalone(t *testing.T) {
	doc := document{
		root: &node{
			nodeType: heading1Node,
			children: []*node{text("Hello")},
		},
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
	}{
		{have: text("Hello <&>"),
			want: "Hello &lt;&amp;&gt;",
		},
		{have: &node{nodeType: horizontalLineNode},
			want: "<hr/>",
		},
		{have: &node{nodeType: lineBreakNode},
			want: "<br/>",
		},
		{have: &node{
			nodeType: italicsNode,
			children: []*node{text("Hello")},
		},
			want: "<i>Hello</i>",
		},
		{have: &node{
			nodeType: boldNode,
			children: []*node{text("Hello")},
		},
			want: "<strong>Hello</strong>",
		},
		{have: &node{
			nodeType: strikeNode,
			children: []*node{text("Hello")},
		},
			want: "<del>Hello</del>",
		},
		{have: &node{
			nodeType: underlinedNode,
			children: []*node{text("Hello")},
		},
			want: "<em>Hello</em>",
		},
		{have: &node{
			nodeType: subscriptNode,
			children: []*node{text("Hello")},
		},
			want: "<sub>Hello</sub>",
		},
		{have: &node{
			nodeType: superscriptNode,
			children: []*node{text("Hello")},
		},
			want: "<sup>Hello</sup>",
		},
		{have: &node{
			nodeType: noWikiNode,
			children: []*node{text("Hello")},
		},
			want: "<pre><tt>Hello</tt></pre>",
		},
		{have: &node{
			nodeType: noWikiInlineNode,
			children: []*node{text("Hello")},
		},
			want: "<tt>Hello</tt>",
		},
		{have: &node{nodeType: heading1Node},
			want: "<h1></h1>",
		},
		{have: &node{nodeType: heading2Node},
			want: "<h2></h2>",
		},
		{have: &node{nodeType: heading3Node},
			want: "<h3></h3>",
		},
		{have: &node{nodeType: heading4Node},
			want: "<h4></h4>",
		},
		{have: &node{nodeType: heading5Node},
			want: "<h5></h5>",
		},
		{have: &node{nodeType: heading6Node},
			want: "<h6></h6>",
		},
		{have: &node{nodeType: paragraphNode},
			want: "<p></p>",
		},
		{have: &node{
			nodeType: linkNode,
			value:    "http://example.com",
			children: []*node{text("Hello")},
		},
			want: `<a href="http://example.com">Hello</a>`,
		},
		{have: &node{
			nodeType: imageNode,
			value: &image{
				src: "http://example.com/image.png",
			},
		},
			want: `<img src="http://example.com/image.png"></img>`,
		},
		{have: &node{
			nodeType: imageNode,
			value: &image{
				src: "http://example.com/image.png",
				alt: "An image",
			},
		},
			want: `<img src="http://example.com/image.png" src="An image"></img>`,
		},
		{have: &node{
			nodeType: unorderedListNode,
			children: []*node{
				&node{
					nodeType: listItemNode,
					children: []*node{text("Hello")},
				},
				&node{
					nodeType: listItemNode,
					children: []*node{text("World")},
				},
			},
		},
			want: `<ul><li>Hello</li><li>World</li></ul>`,
		},
		{have: &node{
			nodeType: orderedListNode,
			children: []*node{
				&node{
					nodeType: listItemNode,
					children: []*node{text("Hello")},
				},
				&node{
					nodeType: listItemNode,
					children: []*node{text("World")},
				},
			},
		},
			want: `<ol><li>Hello</li><li>World</li></ol>`,
		},
		{have: &node{
			nodeType: tableNode,
			children: []*node{
				&node{
					nodeType: tableHeaderRowNode,
					children: []*node{
						&node{
							nodeType: tableHeaderCellNode,
							children: []*node{text("header1")},
						},
						&node{
							nodeType: tableHeaderCellNode,
							children: []*node{text("header2")},
						},
					},
				},
				&node{
					nodeType: tableRowNode,
					children: []*node{
						&node{
							nodeType: tableCellNode,
							children: []*node{text("Hello")},
						},
						&node{
							nodeType: tableCellNode,
							children: []*node{text("World")},
						},
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
		},
	}

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
