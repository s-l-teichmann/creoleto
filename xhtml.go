// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
)

type xhtml struct {
	out *bufio.Writer
}

func (x *xhtml) writeString(s string) error {
	_, err := x.out.WriteString(s)
	return err
}

func (x *xhtml) tag(name string) func(*node) error {
	tag := "<" + name + "/>"
	return func(*node) error { return x.writeString(tag) }
}

func (x *xhtml) open(name string) func(*node) error {
	tag := "<" + name + ">"
	return func(*node) error { return x.writeString(tag) }
}

func (x *xhtml) close(name string) func(*node) error {
	tag := "</" + name + ">"
	return func(*node) error { return x.writeString(tag) }
}

func (x *xhtml) element(name string) *visitor {
	return &visitor{x.open(name), x.close(name)}
}

func (x *xhtml) heading(level int) *visitor {
	tag := fmt.Sprintf("h%d", level)
	return &visitor{x.open(tag), x.close(tag)}
}

func (x *xhtml) text(n *node) error {
	enc := xml.NewEncoder(x.out)
	txt := n.value.(string)
	if err := enc.EncodeToken(xml.CharData(txt)); err != nil {
		return err
	}
	return enc.Flush()
}

func (x *xhtml) noWikiInline(n *node) error {
	x.writeString("<tt>")
	x.text(n)
	return x.writeString("</tt>")
}

func (x *xhtml) noWiki(n *node) error {
	x.writeString("<pre>")
	x.noWikiInline(n)
	return x.writeString("</pre>\n")
}

func (x *xhtml) link(n *node) error {
	href := n.value.(string)
	enc := xml.NewEncoder(x.out)
	if err := enc.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: "a"},
		Attr: []xml.Attr{{xml.Name{Local: "href"}, href}},
	}); err != nil {
		return err
	}
	return enc.Flush()
}

func (x *xhtml) image(n *node) error {
	img := n.value.(*image)
	enc := xml.NewEncoder(x.out)
	attr := []xml.Attr{{xml.Name{Local: "src"}, img.src}}
	if img.alt != "" {
		attr = append(attr, xml.Attr{xml.Name{Local: "src"}, img.alt})
	}
	if err := enc.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: "img"},
		Attr: attr,
	}); err != nil {
		return err
	}
	return enc.Flush()
}

func exportXHTML(doc *document, out io.Writer) error {

	x := xhtml{out: bufio.NewWriter(out)}

	x.out.WriteString(xml.Header)
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
		imageNode:          &visitor{enter: x.image, leave: x.close("img")},
		linkNode:           &visitor{enter: x.link, leave: x.close("a")},
		horizontalLineNode: &visitor{enter: x.tag("hr")},
	})
	if err != nil {
		return err
	}
	x.writeString("\n</body>\n</html>\n")
	return x.out.Flush()
}
