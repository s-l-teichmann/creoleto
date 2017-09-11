// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
// linkChildren repairs the parent/child relations.
package main

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

func nd(typ nodeType, children ...*node) *node {
	return &node{
		nodeType: typ,
		children: children,
	}
}
