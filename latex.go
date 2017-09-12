// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

const (
	latexHeader = `\documentclass[]{article}
\usepackage{lmodern}
\usepackage{amssymb,amsmath}
\usepackage{ifxetex,ifluatex}
\usepackage{cancel}
\usepackage{fixltx2e} % provides \textsubscript
\ifnum 0\ifxetex 1\fi\ifluatex 1\fi=0 % if pdftex
  \usepackage[T1]{fontenc}
  \usepackage[utf8]{inputenc}
\else % if luatex or xelatex
  \ifxetex
    \usepackage{mathspec}
  \else
    \usepackage{fontspec}
  \fi
  \defaultfontfeatures{Ligatures=TeX,Scale=MatchLowercase}
\fi
% use upquote if available, for straight quotes in verbatim environments
\IfFileExists{upquote.sty}{\usepackage{upquote}}{}
% use microtype if available
\IfFileExists{microtype.sty}{%
\usepackage[]{microtype}
\UseMicrotypeSet[protrusion]{basicmath} % disable protrusion for tt fonts
}{}
\PassOptionsToPackage{hyphens}{url} % url is loaded by hyperref
\usepackage[unicode=true]{hyperref}
\hypersetup{
            pdfborder={0 0 0},
            breaklinks=true}
\urlstyle{same}  % don't use monospace font for urls
\IfFileExists{parskip.sty}{%
\usepackage{parskip}
}{% else
\setlength{\parindent}{0pt}
\setlength{\parskip}{6pt plus 2pt minus 1pt}
}
\setlength{\emergencystretch}{3em}  % prevent overfull lines
\providecommand{\tightlist}{%
  \setlength{\itemsep}{0pt}\setlength{\parskip}{0pt}}
\setcounter{secnumdepth}{0}
% Redefines (sub)paragraphs to behave more like sections
\ifx\paragraph\undefined\else
\let\oldparagraph\paragraph
\renewcommand{\paragraph}[1]{\oldparagraph{#1}\mbox{}}
\fi
\ifx\subparagraph\undefined\else
\let\oldsubparagraph\subparagraph
\renewcommand{\subparagraph}[1]{\oldsubparagraph{#1}\mbox{}}
\fi

% set default figure placement to htbp
\makeatletter
\def\fps@figure{htbp}
\makeatother


\date{}

\begin{document}
`
	latexFooter = `
\end{document}
`
)

// latexReplacer substitutes special LaTex characters with quoted ones
// which can be used in templates.
var latexReplacer = strings.NewReplacer(
	"\n", `\\`,
	`\`, `\textbackslash`,
	`$`, `\$`,
	`_`, `\_`,
	`%`, `\%`,
	`&`, `\&`,
	`#`, `\#`,
	`^`, `\url{^}`,
	`~`, `\url{~}`,
	`"`, `\url{"}`,
	`}`, `\}`,
	`{`, `\{`,
	`]`, `{]}`,
	`[`, `{[}`)

type latex struct {
	out    *bufio.Writer
	labels map[string]struct{}
}

func (l *latex) writeString(txt string) error {
	_, err := l.out.WriteString(txt)
	return err
}

func (l *latex) generateLabel(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	name = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '-'
		}
		return r
	}, name)
	if _, found := l.labels[name]; !found {
		l.labels[name] = struct{}{}
		return name
	}
	for i := 1; ; i++ {
		n := name + "-" + strconv.Itoa(i)
		if _, found := l.labels[n]; !found {
			l.labels[n] = struct{}{}
			return n
		}
	}
}

func (l *latex) str(txt string) func(*node) error {
	return func(n *node) error { return l.writeString(txt) }
}

func (l *latex) command(cmd string) *visitor {
	return &visitor{l.str(`\` + cmd + `{`), l.str(`}`)}
}

func (l *latex) text(n *node) error {
	return l.writeString(latexReplacer.Replace(n.value.(string)))
}

func (l *latex) link() *visitor {
	return &visitor{
		enter: func(n *node) error {
			l.writeString(`\href{`)
			l.writeString(latexReplacer.Replace(n.value.(string)))
			return l.writeString(`}{`)
		},
		leave: l.str(`}`),
	}
}

func (l *latex) image(n *node) error {
	img := n.value.(*image)
	l.writeString(`\begin{figure}` + "\n")
	l.writeString(`\centering` + "\n")
	l.writeString(`\includegraphics{`)
	l.writeString(latexReplacer.Replace(img.src))
	l.writeString(`}` + "\n" + `\caption{`)
	l.writeString(latexReplacer.Replace(img.alt))
	return l.writeString(`}` + "\n" + `\end{figure}` + "\n")
}

func (l *latex) header(typ string) func(*node) error {
	return func(n *node) error {
		title := n.value.(string)
		_, err := fmt.Fprintf(l.out, "\n\n\\%s{%s}\\label{%s}\n\n",
			typ,
			latexReplacer.Replace(title),
			latexReplacer.Replace(l.generateLabel(title)))
		return err
	}
}

func (l *latex) table() *visitor {
	return &visitor{
		enter: func(n *node) error {
			align := strings.Repeat("l", n.numColumns())
			return l.writeString(
				"\n\\begin{longtable}[]{@{}" + align + "@{}}\n" +
					"\\toprule\n")
		},
		leave: l.str("\n\\bottomrule\n\\end{longtable}\n"),
	}
}

func (l *latex) tableHeader() *visitor {
	return &visitor{
		leave: func(n *node) error {
			l.writeString("\\tabularnewline\n")
			l.writeString("\\midrule\n")
			return l.writeString("\\endhead\n")
		},
	}
}

func (l *latex) tableCell() *visitor {
	return &visitor{
		leave: func(n *node) error {
			var err error
			if n.parent.nextOfType(n) != nil {
				err = l.writeString(" & ")
			}
			return err
		},
	}
}

func exportLaTex(doc *document, out io.Writer, standalone bool) error {

	l := latex{
		out:    bufio.NewWriter(out),
		labels: make(map[string]struct{}),
	}

	if standalone {
		l.writeString(latexHeader)
	}

	err := doc.traverse(map[nodeType]*visitor{
		orderedListNode: &visitor{
			l.str("\n" + `\begin{enumerate}` + "\n" +
				`\def\labelenumi{\arabic{enumi}.}` + "\n"),
			l.str("\n" + `\end{enumerate}` + "\n")},
		unorderedListNode: &visitor{
			l.str("\n" + `\begin{itemize}` + "\n"),
			l.str("\n" + `\end{itemize}` + "\n")},
		listItemNode: &visitor{
			l.str("\n" + `\item` + "\n"),
			l.str("\n")},
		textNode:            &visitor{enter: l.text},
		boldNode:            l.command("textbf"),
		italicsNode:         l.command("textit"),
		underlinedNode:      l.command("underline"),
		strikeNode:          l.command("cancel"),
		superscriptNode:     l.command("textsuperscript"),
		subscriptNode:       l.command("textsubscript"),
		tableNode:           l.table(),
		tableRowNode:        &visitor{leave: l.str("\\tabularnewline\n")},
		tableCellNode:       l.tableCell(),
		tableHeaderRowNode:  l.tableHeader(),
		tableHeaderCellNode: l.tableCell(),
		heading1Node:        &visitor{enter: l.header("section")},
		heading2Node:        &visitor{enter: l.header("subsection")},
		heading3Node:        &visitor{enter: l.header("subsubsection")},
		heading4Node:        &visitor{enter: l.header("paragraph")},
		heading5Node:        &visitor{enter: l.header("subparagraph")},
		heading6Node:        &visitor{enter: l.header("subparagraph")},
		paragraphNode:       &visitor{l.str("\n"), l.str("\n")},
		lineBreakNode:       &visitor{enter: l.str(`\\` + "\n")},
		escapeNode:          &visitor{enter: l.text},
		noWikiNode: &visitor{
			l.str("\n" + `\begin{verbatim}` + "\n"),
			l.str("\n" + `\end{verbatim}` + "\n")},
		noWikiInlineNode: l.command("texttt"),
		imageNode:        &visitor{enter: l.image},
		linkNode:         l.link(),
		horizontalLineNode: &visitor{
			enter: l.str(
				"\n" + `\begin{center}\rule{0.5\linewidth}{\linethickness}\end{center}` + "\n")},
	})

	if err != nil {
		return err
	}

	if standalone {
		l.writeString(latexFooter)
	}

	return l.out.Flush()
}
