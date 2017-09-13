// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// Copyright 2017 by Intevation GmbH
package main

import "strings"

const (
	proto = `http|https|ftp|nntp|news|mailto|telnet|file|irc`

	//(^ | (?<=\s)) Unsupported perl syntax
	urlRe = `
        (?P<url>
        (^ | (?:\s))
        (?P<escaped_url>~)?
        (?P<url_target> (?P<url_proto> ` + proto + `)://[^$\s]+ )
        )`
	linkRe = `(?P<link>
        \[\[
        (?P<link_target>.+?) \s*
        ([|] \s* (?P<link_text>.+?) \s*)?
        ]]
        )`
	imageRe = `
		(?P<image>
        {{
        (?P<image_target>.+?) \s*
        (\| \s* (?P<image_text>.+?) \s*)?
        }}
        )(?i)`
	// A macro like: <<macro>>text<</macro>>
	// Unsupported Perl syntax <</ \s* (?P=macro_inline_start) \s* >>
	macroInlineRe = `
        (?P<macro_inline>
        << \s* (?P<macro_inline_start>\w+) \s* (?P<macro_inline_args>.*?) \s* >>
        (?P<macro_inline_text>(.|\n)*?)
        <</ \s* (?P<macro_inline_start>) \s* >>
        )`
	// A simple macro tag, like <<macro-a foo="bar">> or <<macro />>
	macroTagRe = `
        (?P<macro_tag>
        <<(?P<macro_tag_name> \w+) (?P<macro_tag_args>.*?) \s* /*>>
        )`

	preInlineRe = `(?P<pre_inline> {{{ (?P<pre_inline_text>.*?) }}} )`

	// Basic text typefaces

	emphasisRe = `(?P<emphasis>(?:|[^:])// (?P<emphasis_text>.+?) (?:|[^:])// )`
	// There must be no ':' in front of the '//' avoid italic
	// rendering in urls with unknown protocols.

	strongRe = `(?P<strong>\*\* (?P<strong_text>.+?) \*\* )`

	// Creole 1.0 optional
	monospaceRe   = `(?P<monospace> \#\# (?P<monospace_text>.+?) \#\# )`
	superscriptRe = `(?P<superscript> \^\^ (?P<superscript_text>.+?) \^\^ )`
	subscriptRe   = `(?P<subscript> ,, (?P<subscript_text>.+?) ,, )`
	underlineRe   = `(?P<underline> __ (?P<underline_text>.+?) __ )`
	deleteRe      = `(?P<delete> ~~ (?P<delete_text>.+?) ~~ )`

	// Own additions
	smallRe = `(?P<small>-- (?P<small_text>.+?) -- )`

	linebreakRe = `(?P<linebreak> \\\\ )`
	escapeRe    = `(?P<escape> ~ (?P<escaped_char>\S) )`
	charRe      = `(?P<char> . )`

	// Block rules
	macroBlockRe = `
        (?P<macro_block>
        << \s* (?P<macro_block_start>\w+) \s* (?P<macro_block_args>.*?) \s* >>
        (?P<macro_block_text>(.|\n)*?)
        <</ \s* (?P<macro_block_start>) \s* >>
        )`

	// Empty line that separates paragraphs.
	lineRe = `(?P<line> ^\s*$ )`

	headRe = `
        (?P<head>
        ^
        (?P<head_head>=+) \s*
        (?P<head_text> .*? )
        (=|\s)*?$
        )`

	// Horizontal line
	separatorRe = `(?P<separator> ^ \s* ---- \s* $ )`

	preBlockRe = `
        (?P<pre_block>
        ^{{{ \s* $
        (?P<pre_block_text>
            ([\#]!(?P<pre_block_kind>\w*?)(\s+.*)?$)?
            (.|\n)+?
        )
        ^}}})`

	// matches thw whole list, separate items are parsed later.
	// The list *must* start with a single bullet.
	listRe = `
        (?P<list>
        ^ [ \t]* ([*][^*\#]|[\#][^\#*]).* $
        ( \n[ \t]* [*\#]+.* $ )*
        )`

	tableRe = `
        ^ \s*(?P<table>
            [|].*? \s*
            [|]?
        ) \s* $`

	// Support blog style tests, too?
	//`(?P<space> (?<!\\)$\n(?!\s*$) )? (?P<text> .+ )`)
	textRe = `(?P<space> \s* )? (?P<text> .+ )`

	itemRe = `
		^ \s* (?P<item>
        (?P<item_head> [\#*]+) \s*
        (?P<item_text> .*?)
        ) \s* $`

	cellRe = `
		\| \s*
            (
                (?P<head> [=][^|]+ ) |
                (?P<cell> (` +
		linkRe + `|` +
		macroInlineRe + `|` +
		macroTagRe + `|` +
		imageRe + `|` +
		preInlineRe + `|` +
		`| [^|])+ )
            ) \s*`

	preEscapeRe = `^(?P<indent>\s*) ~ (?P<rest> \}\}\} \s*) $`
)

var (
	blockRules = []string{
		macroBlockRe,
		lineRe,
		headRe,
		separatorRe,
		preBlockRe,
		listRe,
		tableRe,
		textRe,
	}

	inlineRules = []string{
		linkRe, urlRe,
		macroInlineRe, macroTagRe,
		preInlineRe, imageRe,

		strongRe, emphasisRe,
		monospaceRe, underlineRe,
		superscriptRe, subscriptRe,
		smallRe, deleteRe,

		linebreakRe, escapeRe, charRe,
	}

	blockRE = mustCompileVerbose(`(?m)` +
		strings.Join(blockRules, `|`))
)
