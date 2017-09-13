// Generated from Creole10.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 139,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 5, 7, 74, 10, 7, 3, 7, 3, 7, 5, 7, 78, 10,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 6, 10, 86, 10, 10, 13, 10, 14,
	10, 87, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 2, 2,
	30, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 2, 17, 2, 19, 9, 21, 2, 23,
	2, 25, 10, 27, 11, 29, 12, 31, 13, 33, 14, 35, 15, 37, 16, 39, 17, 41,
	18, 43, 19, 45, 20, 47, 21, 49, 22, 51, 23, 53, 24, 55, 25, 57, 26, 3,
	2, 2, 2, 138, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 3, 59, 3, 2, 2, 2, 5, 61, 3, 2, 2, 2, 7,
	63, 3, 2, 2, 2, 9, 65, 3, 2, 2, 2, 11, 67, 3, 2, 2, 2, 13, 77, 3, 2, 2,
	2, 15, 79, 3, 2, 2, 2, 17, 81, 3, 2, 2, 2, 19, 85, 3, 2, 2, 2, 21, 89,
	3, 2, 2, 2, 23, 91, 3, 2, 2, 2, 25, 93, 3, 2, 2, 2, 27, 96, 3, 2, 2, 2,
	29, 99, 3, 2, 2, 2, 31, 103, 3, 2, 2, 2, 33, 107, 3, 2, 2, 2, 35, 110,
	3, 2, 2, 2, 37, 113, 3, 2, 2, 2, 39, 116, 3, 2, 2, 2, 41, 119, 3, 2, 2,
	2, 43, 122, 3, 2, 2, 2, 45, 124, 3, 2, 2, 2, 47, 126, 3, 2, 2, 2, 49, 128,
	3, 2, 2, 2, 51, 130, 3, 2, 2, 2, 53, 132, 3, 2, 2, 2, 55, 134, 3, 2, 2,
	2, 57, 137, 3, 2, 2, 2, 59, 60, 7, 60, 2, 2, 60, 4, 3, 2, 2, 2, 61, 62,
	7, 69, 2, 2, 62, 6, 3, 2, 2, 2, 63, 64, 7, 52, 2, 2, 64, 8, 3, 2, 2, 2,
	65, 66, 7, 128, 2, 2, 66, 10, 3, 2, 2, 2, 67, 68, 5, 13, 7, 2, 68, 69,
	7, 127, 2, 2, 69, 70, 7, 127, 2, 2, 70, 71, 7, 127, 2, 2, 71, 12, 3, 2,
	2, 2, 72, 74, 5, 15, 8, 2, 73, 72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74,
	75, 3, 2, 2, 2, 75, 78, 5, 17, 9, 2, 76, 78, 5, 15, 8, 2, 77, 73, 3, 2,
	2, 2, 77, 76, 3, 2, 2, 2, 78, 14, 3, 2, 2, 2, 79, 80, 7, 15, 2, 2, 80,
	16, 3, 2, 2, 2, 81, 82, 7, 12, 2, 2, 82, 18, 3, 2, 2, 2, 83, 86, 5, 21,
	11, 2, 84, 86, 5, 23, 12, 2, 85, 83, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 20, 3, 2, 2,
	2, 89, 90, 7, 34, 2, 2, 90, 22, 3, 2, 2, 2, 91, 92, 7, 11, 2, 2, 92, 24,
	3, 2, 2, 2, 93, 94, 7, 60, 2, 2, 94, 95, 7, 49, 2, 2, 95, 26, 3, 2, 2,
	2, 96, 97, 7, 49, 2, 2, 97, 98, 7, 49, 2, 2, 98, 28, 3, 2, 2, 2, 99, 100,
	7, 125, 2, 2, 100, 101, 7, 125, 2, 2, 101, 102, 7, 125, 2, 2, 102, 30,
	3, 2, 2, 2, 103, 104, 7, 127, 2, 2, 104, 105, 7, 127, 2, 2, 105, 106, 7,
	127, 2, 2, 106, 32, 3, 2, 2, 2, 107, 108, 7, 93, 2, 2, 108, 109, 7, 93,
	2, 2, 109, 34, 3, 2, 2, 2, 110, 111, 7, 95, 2, 2, 111, 112, 7, 95, 2, 2,
	112, 36, 3, 2, 2, 2, 113, 114, 7, 125, 2, 2, 114, 115, 7, 125, 2, 2, 115,
	38, 3, 2, 2, 2, 116, 117, 7, 127, 2, 2, 117, 118, 7, 127, 2, 2, 118, 40,
	3, 2, 2, 2, 119, 120, 7, 94, 2, 2, 120, 121, 7, 94, 2, 2, 121, 42, 3, 2,
	2, 2, 122, 123, 7, 63, 2, 2, 123, 44, 3, 2, 2, 2, 124, 125, 7, 126, 2,
	2, 125, 46, 3, 2, 2, 2, 126, 127, 7, 37, 2, 2, 127, 48, 3, 2, 2, 2, 128,
	129, 7, 47, 2, 2, 129, 50, 3, 2, 2, 2, 130, 131, 7, 44, 2, 2, 131, 52,
	3, 2, 2, 2, 132, 133, 7, 49, 2, 2, 133, 54, 3, 2, 2, 2, 134, 135, 7, 66,
	2, 2, 135, 136, 7, 66, 2, 2, 136, 56, 3, 2, 2, 2, 137, 138, 11, 2, 2, 2,
	138, 58, 3, 2, 2, 2, 7, 2, 73, 77, 85, 87, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "'C'", "'2'", "'~'", "", "", "", "", "'//'", "'{{{'", "'}}}'",
	"'[['", "']]'", "'{{'", "'}}'", "'\\\\'", "'='", "'|'", "'#'", "'-'", "'*'",
	"'/'", "'@@'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "ESCAPE", "NOWIKI_BLOCK_CLOSE", "NEWLINE", "BLANKS", "COLON_SLASH",
	"ITAL", "NOWIKI_OPEN", "NOWIKI_CLOSE", "LINK_OPEN", "LINK_CLOSE", "IMAGE_OPEN",
	"IMAGE_CLOSE", "FORCED_LINEBREAK", "EQUAL", "PIPE", "POUND", "DASH", "STAR",
	"SLASH", "EXTENSION", "INSIGNIFICANT_CHAR",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "ESCAPE", "NOWIKI_BLOCK_CLOSE", "NEWLINE", "CR",
	"LF", "BLANKS", "SPACE", "TABULATOR", "COLON_SLASH", "ITAL", "NOWIKI_OPEN",
	"NOWIKI_CLOSE", "LINK_OPEN", "LINK_CLOSE", "IMAGE_OPEN", "IMAGE_CLOSE",
	"FORCED_LINEBREAK", "EQUAL", "PIPE", "POUND", "DASH", "STAR", "SLASH",
	"EXTENSION", "INSIGNIFICANT_CHAR",
}

type Creole10Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCreole10Lexer(input antlr.CharStream) *Creole10Lexer {

	l := new(Creole10Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Creole10.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Creole10Lexer tokens.
const (
	Creole10LexerT__0               = 1
	Creole10LexerT__1               = 2
	Creole10LexerT__2               = 3
	Creole10LexerESCAPE             = 4
	Creole10LexerNOWIKI_BLOCK_CLOSE = 5
	Creole10LexerNEWLINE            = 6
	Creole10LexerBLANKS             = 7
	Creole10LexerCOLON_SLASH        = 8
	Creole10LexerITAL               = 9
	Creole10LexerNOWIKI_OPEN        = 10
	Creole10LexerNOWIKI_CLOSE       = 11
	Creole10LexerLINK_OPEN          = 12
	Creole10LexerLINK_CLOSE         = 13
	Creole10LexerIMAGE_OPEN         = 14
	Creole10LexerIMAGE_CLOSE        = 15
	Creole10LexerFORCED_LINEBREAK   = 16
	Creole10LexerEQUAL              = 17
	Creole10LexerPIPE               = 18
	Creole10LexerPOUND              = 19
	Creole10LexerDASH               = 20
	Creole10LexerSTAR               = 21
	Creole10LexerSLASH              = 22
	Creole10LexerEXTENSION          = 23
	Creole10LexerINSIGNIFICANT_CHAR = 24
)
