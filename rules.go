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

const proto = `http|https|ftp|nntp|news|mailto|telnet|file|irc`

var (
	//(^ | (?<=\s)) Unsupported perl syntax
	urlRe = mustCompileVerbose(`
        (?P<url>
        (^ | (?:\s))
        (?P<escaped_url>~)?
        (?P<url_target> (?P<url_proto> ` + proto + `)://[^$\s]+ )
        )`)
	linkRe = mustCompileVerbose(
		`(?P<link>
        \[\[
        (?P<link_target>.+?) \s*
        ([|] \s* (?P<link_text>.+?) \s*)?
        ]]
        )`)
	imageRe = mustCompileVerbose(`
		(?P<image>
        {{
        (?P<image_target>.+?) \s*
        (\| \s* (?P<image_text>.+?) \s*)?
        }}
        )(?i)`)
	// A macro like: <<macro>>text<</macro>>
	// Unsupported Perl syntax <</ \s* (?P=macro_inline_start) \s* >>
	macroInlineRe = mustCompileVerbose(`
        (?P<macro_inline>
        << \s* (?P<macro_inline_start>\w+) \s* (?P<macro_inline_args>.*?) \s* >>
        (?P<macro_inline_text>(.|\n)*?)
        <</ \s* (?P<macro_inline_start>) \s* >>
        )`)
	// A simple macro tag, like <<macro-a foo="bar">> or <<macro />>
	macroTagRe = mustCompileVerbose(`
        (?P<macro_tag>
        <<(?P<macro_tag_name> \w+) (?P<macro_tag_args>.*?) \s* /*>>
        )`)

	preInlineRe = mustCompileVerbose(
		`(?P<pre_inline> {{{ (?P<pre_inline_text>.*?) }}} )`)

	// Basic text typefaces

	emphasisRe = mustCompileVerbose(
		`(?P<emphasis>(?:|[^:])// (?P<emphasis_text>.+?) (?:|[^:])// )`)
	// There must be no ':' in front of the '//' avoid italic
	// rendering in urls with unknown protocols.

	strongRe = mustCompileVerbose(
		`(?P<strong>\*\* (?P<strong_text>.+?) \*\* )`)

	// Creole 1.0 optional
	monospaceRe = mustCompileVerbose(
		`(?P<monospace> \#\# (?P<monospace_text>.+?) \#\# )`)
	superscriptRe = mustCompileVerbose(
		`(?P<superscript> \^\^ (?P<superscript_text>.+?) \^\^ )`)
	subscriptRe = mustCompileVerbose(
		`(?P<subscript> ,, (?P<subscript_text>.+?) ,, )`)
	underlineRe = mustCompileVerbose(
		`(?P<underline> __ (?P<underline_text>.+?) __ )`)
	deleteRe = mustCompileVerbose(
		`(?P<delete> ~~ (?P<delete_text>.+?) ~~ )`)

	// Own additions
	smallRe = mustCompileVerbose(
		`(?P<small>-- (?P<small_text>.+?) -- )`)

	linebreakRe = mustCompileVerbose(
		`(?P<linebreak> \\\\ )`)
	escapeRe = mustCompileVerbose(
		`(?P<escape> ~ (?P<escaped_char>\S) )`)
	charRe = mustCompileVerbose(`(?P<char> . )`)

	// Block rules
	macroBlockRe = mustCompileVerbose(`
        (?P<macro_block>
        << \s* (?P<macro_block_start>\w+) \s* (?P<macro_block_args>.*?) \s* >>
        (?P<macro_block_text>(.|\n)*?)
        <</ \s* (?P<macro_block_start>) \s* >>
        )`)

	// Empty line that separates paragraphs.
	lineRe = mustCompileVerbose(`(?P<line> ^\s*$ )`)

	headRe = mustCompileVerbose(`
        (?P<head>
        ^
        (?P<head_head>=+) \s*
        (?P<head_text> .*? )
        (=|\s)*?$
        )`)

	// Horizontal line
	separatorRe = mustCompileVerbose(
		`(?P<separator> ^ \s* ---- \s* $ )`)

	preBlockRe = mustCompileVerbose(`
        (?P<pre_block>
        ^{{{ \s* $
        (?P<pre_block_text>
            ([\#]!(?P<pre_block_kind>\w*?)(\s+.*)?$)?
            (.|\n)+?
        )
        ^}}})`)

	// matches thw whole list, separate items are parsed later.
	// The list *must* start with a single bullet.
	listRe = mustCompileVerbose(`
        (?P<list>
        ^ [ \t]* ([*][^*\#]|[\#][^\#*]).* $
        ( \n[ \t]* [*\#]+.* $ )*
        )`)

	tableRe = mustCompileVerbose(`
        ^ \s*(?P<table>
            [|].*? \s*
            [|]?
        ) \s* $`)

	// Support blog style tests, too?
	textRe = mustCompileVerbose(
		//`(?P<space> (?<!\\)$\n(?!\s*$) )? (?P<text> .+ )`)
		`(?P<space> \s* )? (?P<text> .+ )`)

	rules = []*regexp.Regexp{
		macroBlockRe,
		lineRe,
		headRe,
		separatorRe,
		preBlockRe,
		listRe,
		tableRe,
		textRe,
	}

	itemRe = mustCompileVerbose(`
		^ \s* (?P<item>
        (?P<item_head> [\#*]+) \s*
        (?P<item_text> .*?)
        ) \s* $`)

	cellRe = mustCompileVerbose(`
		\| \s*
            (
                (?P<head> [=][^|]+ ) |
                (?P<cell> (` +
		strings.Join([]string{
			linkRe.String(),
			macroInlineRe.String(),
			macroTagRe.String(),
			imageRe.String(),
			preInlineRe.String(),
		}, "|") +
		`| [^|])+ )
            ) \s*`)

	preEscapeRe = mustCompileVerbose(
		`^(?P<indent>\s*) ~ (?P<rest> \}\}\} \s*) $`)

	inlineRules = []*regexp.Regexp{
		linkRe, urlRe,
		macroInlineRe, macroTagRe,
		preInlineRe, imageRe,

		strongRe, emphasisRe,
		monospaceRe, underlineRe,
		superscriptRe, subscriptRe,
		smallRe, deleteRe,

		linebreakRe, escapeRe, charRe,
	}
)

var ()
