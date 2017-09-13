// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import "strings"

type parser struct {
}

func newParser() *parser {
	return &parser{}
}

func preprocess(s string) string {
	s = strings.Replace(s, "\r\n", "\n", -1)
	s = strings.Replace(s, "\r", "\n", -1)
	return s
}

func (p *parser) parse(input string) (*document, error) {
	input = preprocess(input)
	return nil, nil
}
