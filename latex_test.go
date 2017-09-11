// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import "testing"

func TestLaTexElements(t *testing.T) {
	// TODO: Implement me!
	cases := []testCase{{
		have: nd(noWikiInlineNode, text("Hello")),
		want: `\texttt{Hello}`,
	}, {
		have: text(`^`),
		want: `\url{^}`,
	}}

	runCases(cases, exportLaTex, t)
}
