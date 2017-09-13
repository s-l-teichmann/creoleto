// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"regexp"
	"strings"
	"unicode"
)

func mustCompileVerbose(str string) *regexp.Regexp {
	str = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
	return regexp.MustCompile(str)
}

type matchResult struct {
	parts []string
	re    *regexp.Regexp
}

func match(re *regexp.Regexp, s string) []*matchResult {
	m := re.FindAllStringSubmatch(s, -1)
	if len(m) == 0 {
		return nil
	}
	res := make([]*matchResult, len(m))
	for i := range res {
		res[i] = &matchResult{parts: m[i], re: re}
	}
	return res
}

func (m *matchResult) group(name string) string {
	for i, gname := range m.re.SubexpNames() {
		if gname == name {
			return m.parts[i]
		}
	}
	return ""
}

func (m *matchResult) found(fn func(key, value string)) {
	names := m.re.SubexpNames()
	for i, p := range m.parts {
		if p != "" && names[i] != "" {
			fn(names[i], p)
		}
	}
}
