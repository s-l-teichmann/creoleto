// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"bytes"
	"fmt"
	"io"
)

type xhtml struct {
	out bytes.Buffer
}

func (x *xhtml) italicsEnter(*node) error {
	x.out.WriteString("<i>")
	return nil
}

func (x *xhtml) italicsLeave(*node) error {
	x.out.WriteString("</i>")
	return nil
}

func (x *xhtml) boldEnter(*node) error {
	x.out.WriteString("<strong>")
	return nil
}

func (x *xhtml) boldLeave(*node) error {
	x.out.WriteString("</strong>")
	return nil
}

func (x *xhtml) horizontalLine(*node) error {
	x.out.WriteString("<hr/>")
	return nil
}

func (x *xhtml) heading(level int) *visitor {
	return &visitor{
		enter: func(*node) error {
			fmt.Fprintf(&x.out, "<h%d>", level)
			return nil
		},
		leave: func(*node) error {
			fmt.Fprintf(&x.out, "</h%d>", level)
			return nil
		},
	}
}

func exportXHTML(doc *document, out io.Writer) error {
	// TODO: Implement me!

	var x xhtml
	err := doc.traverse(map[nodeType]*visitor{
		italicsNode:        &visitor{x.italicsEnter, x.italicsLeave},
		boldNode:           &visitor{x.boldEnter, x.boldLeave},
		horizontalLineNode: &visitor{leave: x.horizontalLine},
		heading1Node:       x.heading(1),
		heading2Node:       x.heading(2),
		heading3Node:       x.heading(3),
		heading4Node:       x.heading(4),
		heading5Node:       x.heading(5),
		heading6Node:       x.heading(6),
	})
	if err != nil {
		return err
	}
	_, err = x.out.WriteTo(out)
	return err
}
