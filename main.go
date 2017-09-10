// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func main() {
	format := flag.String(
		"format", "xhtml", `Format to export to ("xhtml" or "latex").`)
	standalone := flag.Bool(
		"standalone", false, `Generate a complete standalone document.`)
	flag.Parse()

	var export func(*document, io.Writer, bool) error
	switch strings.ToLower(*format) {
	case "xhtml":
		export = exportXHTML
	case "latex":
		export = exportLaTex
	default:
		log.Fatalln("Unknown export format.")
	}

	data, err := ioutil.ReadAll(os.Stdin)
	check(err)
	input := string(data)
	p := newParser()
	doc, err := p.parse(input)
	check(err)
	check(export(doc, os.Stdout, *standalone))
}
