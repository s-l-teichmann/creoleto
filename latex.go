// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import (
	"bufio"
	"io"
	"strings"
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
	out *bufio.Writer
}

func (l *latex) writeString(txt string) error {
	_, err := l.out.WriteString(txt)
	return err
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

func exportLaTex(doc *document, out io.Writer, standalone bool) error {

	l := latex{out: bufio.NewWriter(out)}

	if standalone {
		l.writeString(latexHeader)
	}

	// TODO: Implement document traversal.
	err := doc.traverse(map[nodeType]*visitor{
		// TODO: orderedListNode
		// TODO: unorderedListNode
		// TODO: listItemNode
		textNode:        &visitor{enter: l.text},
		boldNode:        l.command("textbf"),
		italicsNode:     l.command("textit"),
		underlinedNode:  l.command("underline"),
		strikeNode:      l.command("cancel"),
		superscriptNode: l.command("textsuperscript"),
		subscriptNode:   l.command("textsubscript"),
		// TODO: tableNode
		// TODO: tableRowNode
		// TODO: tableCellNode
		// TODO: tableHeaderRowNode
		// TODO: tableHeaderCellNode
		// TODO: heading1Node
		// TODO: heading2Node
		// TODO: heading3Node
		// TODO: heading4Node
		// TODO: heading5Node
		// TODO: heading6Node
		paragraphNode: &visitor{l.str("\n"), l.str("\n")},
		lineBreakNode: &visitor{enter: l.str(`\\` + "\n")},
		// TODO: escapeNode
		// TODO: noWikiNode
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
