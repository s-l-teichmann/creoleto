// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

type nodeType int

const (
	rootNode nodeType = iota
	orderedListNode
	unorderedListNode
	listItemNode
	textNode
	boldNode
	italicsNode
	underlinedNode
	strikeNode
	superscriptNode
	subscriptNode
	tableNode
	tableHeaderRowNode
	tableHeaderCellNode
	tableRowNode
	tableCellNode
	heading1Node
	heading2Node
	heading3Node
	heading4Node
	heading5Node
	heading6Node
	paragraphNode
	lineBreakNode
	escapeNode
	noWikiNode
	noWikiInlineNode
	placeholderNode
	imageNode
	linkNode
	horizontalLineNode
)

type node struct {
	nodeType

	parent   *node
	children []*node
	value    interface{}
}

type image struct {
	src string
	alt string
}

func (n *node) nextOfType(c *node) *node {
	if n == nil {
		return nil
	}
	var found bool
	for _, x := range n.children {
		if found {
			if x.nodeType == c.nodeType {
				return x
			}
		} else {
			found = x == c
		}
	}
	return nil
}

func (n *node) numColumns() int {
	var max int
	var descend func(*node)
	descend = func(current *node) {
		if current == nil {
			return
		}
		if current.nodeType == tableRowNode ||
			current.nodeType == tableHeaderRowNode {
			var count int
			for _, c := range current.children {
				if c.nodeType == tableCellNode ||
					c.nodeType == tableHeaderCellNode {
					count++
				}
			}
			if count > max {
				max = count
			}
			return
		}
		for _, c := range current.children {
			descend(c)
		}
	}
	descend(n)
	return max
}

type document struct {
	root *node
}

type visitor struct {
	enter func(*node) error
	leave func(*node) error
}

func (d *document) traverse(visitors map[nodeType]*visitor) error {

	if d == nil {
		return nil
	}

	var descend func(*node) error
	descend = func(n *node) error {
		if n == nil {
			return nil
		}
		v := visitors[n.nodeType]
		if v != nil && v.enter != nil {
			if err := v.enter(n); err != nil {
				return err
			}
		}

		for _, child := range n.children {
			if err := descend(child); err != nil {
				return err
			}
		}

		if v != nil && v.leave != nil {
			if err := v.leave(n); err != nil {
				return err
			}
		}
		return nil
	}

	return descend(d.root)
}

func text(txt string) *node {
	return &node{
		nodeType: textNode,
		value:    txt,
	}
}

func nd(typ nodeType, children ...*node) *node {
	return &node{
		nodeType: typ,
		children: children,
	}
}

func ndp(typ nodeType, parent *node, children ...*node) *node {
	n := nd(typ, children...)
	parent.children = append(parent.children, n)
	return n
}
