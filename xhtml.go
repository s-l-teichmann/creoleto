// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

var htmlEscape = strings.NewReplacer(
	"<", "&lt;",
	">", "&gt;",
	"&", "&amp;").Replace

type xhtml struct {
	out bytes.Buffer
}

func (x *xhtml) tag(name string) func(*node) error {
	tag := "<" + name + "/>"
	return func(*node) error {
		x.out.WriteString(tag)
		return nil
	}
}

func (x *xhtml) open(name string) func(*node) error {
	tag := "<" + name + ">"
	return func(*node) error {
		x.out.WriteString(tag)
		return nil
	}
}

func (x *xhtml) close(name string) func(*node) error {
	tag := "</" + name + ">"
	return func(*node) error {
		x.out.WriteString(tag)
		return nil
	}
}

func (x *xhtml) element(name string) *visitor {
	return &visitor{x.open(name), x.close(name)}
}

func (x *xhtml) heading(level int) *visitor {
	tag := fmt.Sprintf("h%d", level)
	return &visitor{x.open(tag), x.close(tag)}
}

func (x *xhtml) text(n *node) error {
	x.out.WriteString(htmlEscape(n.value.(string)))
	return nil
}

func (x *xhtml) noWikiInline(n *node) error {
	x.out.WriteString("<tt>")
	x.out.WriteString(htmlEscape(n.value.(string)))
	x.out.WriteString("</tt>")
	return nil
}

func (x *xhtml) noWiki(n *node) error {
	x.out.WriteString("<pre>")
	x.noWikiInline(n)
	x.out.WriteString("</pre>\n")
	return nil
}

func (x *xhtml) link(n *node) error {
	href := n.value.(string)
	enc := xml.NewEncoder(&x.out)
	enc.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: "a"},
		Attr: []xml.Attr{{xml.Name{Local: "href"}, href}},
	})
	enc.Flush()
	return nil
}

func exportXHTML(doc *document, out io.Writer) error {
	// TODO: Implement me!

	var x xhtml

	x.out.WriteString("<html>\n<body>\n")

	err := doc.traverse(map[nodeType]*visitor{
		orderedListNode:     x.element("ol"),
		unorderedListNode:   x.element("ul"),
		listItemNode:        x.element("li"),
		textNode:            &visitor{enter: x.text},
		boldNode:            x.element("strong"),
		italicsNode:         x.element("i"),
		underlinedNode:      x.element("em"),
		strikeNode:          x.element("del"),
		superscriptNode:     x.element("sup"),
		subscriptNode:       x.element("sub"),
		tableNode:           x.element("table"),
		tableRowNode:        x.element("tr"),
		tableCellNode:       x.element("td"),
		tableHeaderRowNode:  x.element("th"),
		tableHeaderCellNode: x.element("td"),
		heading1Node:        x.heading(1),
		heading2Node:        x.heading(2),
		heading3Node:        x.heading(3),
		heading4Node:        x.heading(4),
		heading5Node:        x.heading(5),
		heading6Node:        x.heading(6),
		paragraphNode:       x.element("p"),
		lineBreakNode:       &visitor{enter: x.tag("br")},
		escapeNode:          &visitor{enter: x.text},
		noWikiNode:          &visitor{enter: x.noWiki},
		noWikiInlineNode:    &visitor{enter: x.noWikiInline},
		// placeholderNode not supported, yet.
		// TODO: Implement image node.
		linkNode:           &visitor{enter: x.link, leave: x.close("a")},
		horizontalLineNode: &visitor{enter: x.tag("hr")},
	})
	if err != nil {
		x.out.WriteString("\n</body>\n</html>\n")
		_, err = x.out.WriteTo(out)
	}
	return err
}
