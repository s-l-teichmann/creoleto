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
	}, {
		have: &node{
			nodeType: imageNode,
			value: &image{
				src: "http://www.example.org/sample.png",
				alt: "A sample image",
			}},
		want: `\begin{figure}` + "\n" +
			`\centering` + "\n" +
			`\includegraphics{http://www.example.org/sample.png}` + "\n" +
			`\caption{A sample image}` + "\n" +
			`\end{figure}` + "\n",
	}, {
		have: &node{
			nodeType: heading1Node,
			value:    "Header 1"},
		want: "\n\n\\section{Header 1}\\label{header-1}\n\n",
	}, {
		have: &node{
			nodeType: heading2Node,
			value:    "Header 2"},
		want: "\n\n\\subsection{Header 2}\\label{header-2}\n\n",
	}, {
		have: &node{
			nodeType: heading3Node,
			value:    "Header 3"},
		want: "\n\n\\subsubsection{Header 3}\\label{header-3}\n\n",
	}, {
		have: &node{
			nodeType: heading4Node,
			value:    "Header 4"},
		want: "\n\n\\paragraph{Header 4}\\label{header-4}\n\n",
	}, {
		have: &node{
			nodeType: heading5Node,
			value:    "Header 5"},
		want: "\n\n\\subparagraph{Header 5}\\label{header-5}\n\n",
	}, {
		have: &node{
			nodeType: heading6Node,
			value:    "Header 6"},
		want: "\n\n\\subparagraph{Header 6}\\label{header-6}\n\n",
	}, {
		have: nd(noWikiNode, text(`\end{verbatim}`)),
		want: "\n" + `\begin{verbatim}` + "\n" +
			`\textbackslashend\{verbatim\}` + "\n" +
			"\\end{verbatim}\n",
	}, {
		have: &node{
			nodeType: escapeNode,
			value:    "Some text"},
		want: "Some text",
	}, {
		have: nd(unorderedListNode,
			nd(listItemNode, text("First")),
			nd(listItemNode, text("Second"))),
		want: "\n\\begin{itemize}\n\n" +
			"\\item\n" +
			"First\n\n" +
			"\\item\n" +
			"Second\n\n" +
			"\\end{itemize}\n",
	}, {
		have: nd(orderedListNode,
			nd(listItemNode, text("First")),
			nd(listItemNode, text("Second"))),
		want: "\n\\begin{enumerate}\n" +
			`\def\labelenumi{\arabic{enumi}.}` + "\n\n" +
			"\\item\n" +
			"First\n\n" +
			"\\item\n" +
			"Second\n\n" +
			"\\end{enumerate}\n",
	}, {
		have: nd(tableNode,
			nd(tableHeaderRowNode,
				nd(tableHeaderCellNode, text("A")),
				nd(tableHeaderCellNode, text("B")),
				nd(tableHeaderCellNode, text("C"))),
			nd(tableRowNode,
				nd(tableCellNode, text("a")),
				nd(tableCellNode, text("b")),
				nd(tableCellNode, text("c"))),
			nd(tableRowNode,
				nd(tableCellNode, text("1")),
				nd(tableCellNode, text("2")),
				nd(tableCellNode, text("3")))),
		want: "\n" +
			`\begin{longtable}[]{@{}lll@{}}` + "\n" +
			`\toprule` + "\n" +
			`A & B & C\tabularnewline` + "\n" +
			`\midrule` + "\n" +
			`\endhead` + "\n" +
			`a & b & c\tabularnewline` + "\n" +
			`1 & 2 & 3\tabularnewline` + "\n\n" +
			`\bottomrule` + "\n" +
			`\end{longtable}` + "\n",
	}}

	runCases(cases, exportLaTex, t)
}
