// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

type nodeType int

type node struct {
	nodeType

	parent   *node
	children []*node
}

type document struct {
	root *node
}
