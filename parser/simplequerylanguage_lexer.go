// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SimpleQueryLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var simplequerylanguagelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplequerylanguagelexerLexerInit() {
	staticData := &simplequerylanguagelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "':'", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='", "'&&'", "'||'",
		"'!'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['", "']'", "'?'",
		"'??'", "'.'", "','", "'{'", "'}'", "'null'", "'false'", "'true'",
	}
	staticData.symbolicNames = []string{
		"", "", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Question",
		"NullCoalescing", "Dot", "Comma", "LeftBrace", "RightBrace", "Null",
		"False", "True", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual",
		"NotEqual", "And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "Question",
		"NullCoalescing", "Dot", "Comma", "LeftBrace", "RightBrace", "Null",
		"False", "True", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 180, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 139,
		8, 27, 10, 27, 12, 27, 142, 9, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5,
		27, 149, 8, 27, 10, 27, 12, 27, 152, 9, 27, 1, 27, 3, 27, 155, 8, 27, 1,
		28, 1, 28, 5, 28, 159, 8, 28, 10, 28, 12, 28, 162, 9, 28, 1, 29, 4, 29,
		165, 8, 29, 11, 29, 12, 29, 166, 1, 29, 1, 29, 4, 29, 171, 8, 29, 11, 29,
		12, 29, 172, 3, 29, 175, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 0, 0, 31, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 1, 0, 6, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 4,
		0, 36, 36, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 188, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 63, 1, 0,
		0, 0, 3, 65, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 70, 1, 0, 0, 0, 9, 73, 1,
		0, 0, 0, 11, 75, 1, 0, 0, 0, 13, 78, 1, 0, 0, 0, 15, 81, 1, 0, 0, 0, 17,
		84, 1, 0, 0, 0, 19, 87, 1, 0, 0, 0, 21, 89, 1, 0, 0, 0, 23, 91, 1, 0, 0,
		0, 25, 93, 1, 0, 0, 0, 27, 95, 1, 0, 0, 0, 29, 97, 1, 0, 0, 0, 31, 99,
		1, 0, 0, 0, 33, 101, 1, 0, 0, 0, 35, 103, 1, 0, 0, 0, 37, 105, 1, 0, 0,
		0, 39, 107, 1, 0, 0, 0, 41, 110, 1, 0, 0, 0, 43, 112, 1, 0, 0, 0, 45, 114,
		1, 0, 0, 0, 47, 116, 1, 0, 0, 0, 49, 118, 1, 0, 0, 0, 51, 123, 1, 0, 0,
		0, 53, 129, 1, 0, 0, 0, 55, 154, 1, 0, 0, 0, 57, 156, 1, 0, 0, 0, 59, 164,
		1, 0, 0, 0, 61, 176, 1, 0, 0, 0, 63, 64, 5, 58, 0, 0, 64, 2, 1, 0, 0, 0,
		65, 66, 5, 60, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 5, 60, 0, 0, 68, 69, 5,
		61, 0, 0, 69, 6, 1, 0, 0, 0, 70, 71, 5, 61, 0, 0, 71, 72, 5, 61, 0, 0,
		72, 8, 1, 0, 0, 0, 73, 74, 5, 62, 0, 0, 74, 10, 1, 0, 0, 0, 75, 76, 5,
		62, 0, 0, 76, 77, 5, 61, 0, 0, 77, 12, 1, 0, 0, 0, 78, 79, 5, 33, 0, 0,
		79, 80, 5, 61, 0, 0, 80, 14, 1, 0, 0, 0, 81, 82, 5, 38, 0, 0, 82, 83, 5,
		38, 0, 0, 83, 16, 1, 0, 0, 0, 84, 85, 5, 124, 0, 0, 85, 86, 5, 124, 0,
		0, 86, 18, 1, 0, 0, 0, 87, 88, 5, 33, 0, 0, 88, 20, 1, 0, 0, 0, 89, 90,
		5, 43, 0, 0, 90, 22, 1, 0, 0, 0, 91, 92, 5, 45, 0, 0, 92, 24, 1, 0, 0,
		0, 93, 94, 5, 42, 0, 0, 94, 26, 1, 0, 0, 0, 95, 96, 5, 47, 0, 0, 96, 28,
		1, 0, 0, 0, 97, 98, 5, 40, 0, 0, 98, 30, 1, 0, 0, 0, 99, 100, 5, 41, 0,
		0, 100, 32, 1, 0, 0, 0, 101, 102, 5, 91, 0, 0, 102, 34, 1, 0, 0, 0, 103,
		104, 5, 93, 0, 0, 104, 36, 1, 0, 0, 0, 105, 106, 5, 63, 0, 0, 106, 38,
		1, 0, 0, 0, 107, 108, 5, 63, 0, 0, 108, 109, 5, 63, 0, 0, 109, 40, 1, 0,
		0, 0, 110, 111, 5, 46, 0, 0, 111, 42, 1, 0, 0, 0, 112, 113, 5, 44, 0, 0,
		113, 44, 1, 0, 0, 0, 114, 115, 5, 123, 0, 0, 115, 46, 1, 0, 0, 0, 116,
		117, 5, 125, 0, 0, 117, 48, 1, 0, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120,
		5, 117, 0, 0, 120, 121, 5, 108, 0, 0, 121, 122, 5, 108, 0, 0, 122, 50,
		1, 0, 0, 0, 123, 124, 5, 102, 0, 0, 124, 125, 5, 97, 0, 0, 125, 126, 5,
		108, 0, 0, 126, 127, 5, 115, 0, 0, 127, 128, 5, 101, 0, 0, 128, 52, 1,
		0, 0, 0, 129, 130, 5, 116, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 117,
		0, 0, 132, 133, 5, 101, 0, 0, 133, 54, 1, 0, 0, 0, 134, 140, 5, 34, 0,
		0, 135, 136, 5, 92, 0, 0, 136, 139, 9, 0, 0, 0, 137, 139, 8, 0, 0, 0, 138,
		135, 1, 0, 0, 0, 138, 137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138,
		1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 143, 155, 5, 34, 0, 0, 144, 150, 5, 39, 0, 0, 145, 146, 5, 92, 0,
		0, 146, 149, 9, 0, 0, 0, 147, 149, 8, 1, 0, 0, 148, 145, 1, 0, 0, 0, 148,
		147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 155, 5, 39,
		0, 0, 154, 134, 1, 0, 0, 0, 154, 144, 1, 0, 0, 0, 155, 56, 1, 0, 0, 0,
		156, 160, 7, 2, 0, 0, 157, 159, 7, 3, 0, 0, 158, 157, 1, 0, 0, 0, 159,
		162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 58, 1,
		0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 165, 7, 4, 0, 0, 164, 163, 1, 0, 0,
		0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		174, 1, 0, 0, 0, 168, 170, 3, 41, 20, 0, 169, 171, 7, 4, 0, 0, 170, 169,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0,
		0, 0, 173, 175, 1, 0, 0, 0, 174, 168, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0,
		175, 60, 1, 0, 0, 0, 176, 177, 7, 5, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179,
		6, 30, 0, 0, 179, 62, 1, 0, 0, 0, 10, 0, 138, 140, 148, 150, 154, 160,
		166, 172, 174, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SimpleQueryLanguageLexerInit initializes any static state used to implement SimpleQueryLanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleQueryLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleQueryLanguageLexerInit() {
	staticData := &simplequerylanguagelexerLexerStaticData
	staticData.once.Do(simplequerylanguagelexerLexerInit)
}

// NewSimpleQueryLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleQueryLanguageLexer(input antlr.CharStream) *SimpleQueryLanguageLexer {
	SimpleQueryLanguageLexerInit()
	l := new(SimpleQueryLanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &simplequerylanguagelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SimpleQueryLanguage.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleQueryLanguageLexer tokens.
const (
	SimpleQueryLanguageLexerT__0           = 1
	SimpleQueryLanguageLexerLess           = 2
	SimpleQueryLanguageLexerLessOrEqual    = 3
	SimpleQueryLanguageLexerEqual          = 4
	SimpleQueryLanguageLexerGreater        = 5
	SimpleQueryLanguageLexerGreaterOrEqual = 6
	SimpleQueryLanguageLexerNotEqual       = 7
	SimpleQueryLanguageLexerAnd            = 8
	SimpleQueryLanguageLexerOr             = 9
	SimpleQueryLanguageLexerNot            = 10
	SimpleQueryLanguageLexerAdd            = 11
	SimpleQueryLanguageLexerSubtract       = 12
	SimpleQueryLanguageLexerMultiply       = 13
	SimpleQueryLanguageLexerDivide         = 14
	SimpleQueryLanguageLexerLeftParen      = 15
	SimpleQueryLanguageLexerRightParen     = 16
	SimpleQueryLanguageLexerLeftBracket    = 17
	SimpleQueryLanguageLexerRightBracket   = 18
	SimpleQueryLanguageLexerQuestion       = 19
	SimpleQueryLanguageLexerNullCoalescing = 20
	SimpleQueryLanguageLexerDot            = 21
	SimpleQueryLanguageLexerComma          = 22
	SimpleQueryLanguageLexerLeftBrace      = 23
	SimpleQueryLanguageLexerRightBrace     = 24
	SimpleQueryLanguageLexerNull           = 25
	SimpleQueryLanguageLexerFalse          = 26
	SimpleQueryLanguageLexerTrue           = 27
	SimpleQueryLanguageLexerESCAPED_STRING = 28
	SimpleQueryLanguageLexerSTRING         = 29
	SimpleQueryLanguageLexerNUMBER         = 30
	SimpleQueryLanguageLexerWS             = 31
)
