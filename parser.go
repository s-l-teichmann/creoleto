// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"fmt"
	"strings"
)

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
	p.parseBlock(input)
	return nil, nil
}

func (p *parser) parseBlock(input string) {
	res := match(blockRE, input)
	for _, r := range res {
		r.found(func(k, v string) {
			fmt.Printf("\t%s: '%s'\n", k, v)
		})
		fmt.Println("-------------")
	}

}
