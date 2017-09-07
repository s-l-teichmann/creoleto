// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

type nodeType int

const (
	orderedListNode nodeType = iota
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

	var decent func(*node) error
	decent = func(n *node) error {
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
			if err := decent(child); err != nil {
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

	return decent(d.root)
}
