// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"io/ioutil"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	check(err)
	input := string(data)
	p := newParser()
	d, err := p.parse(input)
	check(err)
	_ = d
}
