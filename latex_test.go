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
	}, {
		have: nd(italicsNode, text("Hello")),
		want: `\textit{Hello}`,
	}, {
		have: nd(underlinedNode, text("Hello")),
		want: `\underline{Hello}`,
	}, {
		have: nd(subscriptNode, text("Hello")),
		want: `\textsubscript{Hello}`,
	}, {
		have: nd(superscriptNode, text("Hello")),
		want: `\textsuperscript{Hello}`,
	}, {
		have: nd(strikeNode, text("Hello")),
		want: `\cancel{Hello}`,
	}, {
		have: &node{nodeType: horizontalLineNode},
		want: "\n" + `\begin{center}\rule{0.5\linewidth}{\linethickness}\end{center}` + "\n",
	}, {
		have: nd(paragraphNode, text("Hello")),
		want: "\nHello\n",
	}, {
		have: &node{
			nodeType: linkNode,
			value:    "http://www.example.org",
			children: []*node{text("Hello")}},
		want: `\href{http://www.example.org}{Hello}`,
	}, {
		have: nd(boldNode, text("Hello")),
		want: `\textbf{Hello}`,
	}, {
		have: &node{nodeType: lineBreakNode},
		want: `\\` + "\n",
	}}

	runCases(cases, exportLaTex, t)
}
