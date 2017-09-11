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

func exportLaTex(doc *document, out io.Writer, standalone bool) error {

	l := latex{out: bufio.NewWriter(out)}

	if standalone {
		l.writeString(latexHeader)
	}

	// TODO: Implement document traversal.
	err := doc.traverse(map[nodeType]*visitor{
		textNode:         &visitor{enter: l.text},
		noWikiInlineNode: l.command("texttt"),
		boldNode:         l.command("textbf"),
		italicsNode:      l.command("textit"),
		underlinedNode:   l.command("underline"),
		strikeNode:       l.command("cancel"),
	})

	if err != nil {
		return err
	}

	if standalone {
		l.writeString(latexFooter)
	}

	return l.out.Flush()
}
