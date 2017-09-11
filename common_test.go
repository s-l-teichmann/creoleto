// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
// linkChildren repairs the parent/child relations.
package main

import (
	"bytes"
	"testing"
)

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

type testCase struct {
	have *node
	want string
}

func runCases(cases []testCase, t *testing.T) {
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
