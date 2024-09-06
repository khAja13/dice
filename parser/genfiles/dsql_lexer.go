// Code generated from DSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DSQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DSQLLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dsqllexerLexerInit() {
	staticData := &DSQLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "';'", "','", "'('", "')'", "'-'", "'+'", "'~'", "'||'", "'*'",
		"'/'", "'%'", "'<<'", "'>>'", "'&'", "'|'", "'<'", "'<='", "'>'", "'>='",
		"'='", "'=='", "'!='", "'<>'", "'$key'", "'$value'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "KEY_TOKEN", "VALUE_TOKEN", "K_SELECT",
		"K_FROM", "K_WHERE", "K_ORDER", "K_BY", "K_LIMIT", "K_ASC", "K_DESC",
		"K_IS", "K_NOT", "K_NULL", "K_LIKE", "K_AND", "K_OR", "VALUE_TOKEN_TRAILING_DOT",
		"IDENTIFIER", "NUMERIC_LITERAL", "STRING_LITERAL", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "KEY_TOKEN", "VALUE_TOKEN",
		"K_SELECT", "K_FROM", "K_WHERE", "K_ORDER", "K_BY", "K_LIMIT", "K_ASC",
		"K_DESC", "K_IS", "K_NOT", "K_NULL", "K_LIKE", "K_AND", "K_OR", "VALUE_TOKEN_TRAILING_DOT",
		"IDENTIFIER", "IDENTIFIER_CHAR", "NUMERIC_LITERAL", "STRING_LITERAL",
		"WHITESPACE", "DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 384, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40, 284, 8, 40, 10,
		40, 12, 40, 287, 9, 40, 1, 41, 1, 41, 1, 42, 4, 42, 292, 8, 42, 11, 42,
		12, 42, 293, 1, 42, 1, 42, 5, 42, 298, 8, 42, 10, 42, 12, 42, 301, 9, 42,
		3, 42, 303, 8, 42, 1, 42, 1, 42, 4, 42, 307, 8, 42, 11, 42, 12, 42, 308,
		3, 42, 311, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 317, 8, 43, 10, 43,
		12, 43, 320, 9, 43, 1, 43, 1, 43, 1, 44, 4, 44, 325, 8, 44, 11, 44, 12,
		44, 326, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48,
		1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58,
		1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1,
		64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69,
		1, 69, 1, 70, 1, 70, 1, 71, 1, 71, 0, 0, 72, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 0, 85, 42, 87, 43, 89, 44, 91, 0, 93, 0, 95, 0, 97, 0, 99, 0, 101,
		0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119,
		0, 121, 0, 123, 0, 125, 0, 127, 0, 129, 0, 131, 0, 133, 0, 135, 0, 137,
		0, 139, 0, 141, 0, 143, 0, 1, 0, 32, 4, 0, 36, 36, 65, 90, 95, 95, 97,
		122, 5, 0, 40, 41, 91, 91, 93, 93, 123, 123, 125, 125, 9, 0, 33, 33, 35,
		38, 42, 43, 45, 58, 60, 90, 94, 95, 97, 122, 124, 124, 126, 126, 1, 0,
		39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 2, 0, 65, 65, 97, 97,
		2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100, 2,
		0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103, 103, 2,
		0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106, 106, 2,
		0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109, 109, 2,
		0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112, 112, 2,
		0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115, 115, 2,
		0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118, 118, 2,
		0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121, 121, 2,
		0, 90, 90, 122, 122, 365, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 145, 1,
		0, 0, 0, 3, 147, 1, 0, 0, 0, 5, 149, 1, 0, 0, 0, 7, 151, 1, 0, 0, 0, 9,
		153, 1, 0, 0, 0, 11, 155, 1, 0, 0, 0, 13, 157, 1, 0, 0, 0, 15, 159, 1,
		0, 0, 0, 17, 162, 1, 0, 0, 0, 19, 164, 1, 0, 0, 0, 21, 166, 1, 0, 0, 0,
		23, 168, 1, 0, 0, 0, 25, 171, 1, 0, 0, 0, 27, 174, 1, 0, 0, 0, 29, 176,
		1, 0, 0, 0, 31, 178, 1, 0, 0, 0, 33, 180, 1, 0, 0, 0, 35, 183, 1, 0, 0,
		0, 37, 185, 1, 0, 0, 0, 39, 188, 1, 0, 0, 0, 41, 190, 1, 0, 0, 0, 43, 193,
		1, 0, 0, 0, 45, 196, 1, 0, 0, 0, 47, 199, 1, 0, 0, 0, 49, 204, 1, 0, 0,
		0, 51, 211, 1, 0, 0, 0, 53, 218, 1, 0, 0, 0, 55, 223, 1, 0, 0, 0, 57, 229,
		1, 0, 0, 0, 59, 235, 1, 0, 0, 0, 61, 238, 1, 0, 0, 0, 63, 244, 1, 0, 0,
		0, 65, 248, 1, 0, 0, 0, 67, 253, 1, 0, 0, 0, 69, 256, 1, 0, 0, 0, 71, 260,
		1, 0, 0, 0, 73, 265, 1, 0, 0, 0, 75, 270, 1, 0, 0, 0, 77, 274, 1, 0, 0,
		0, 79, 277, 1, 0, 0, 0, 81, 280, 1, 0, 0, 0, 83, 288, 1, 0, 0, 0, 85, 310,
		1, 0, 0, 0, 87, 312, 1, 0, 0, 0, 89, 324, 1, 0, 0, 0, 91, 330, 1, 0, 0,
		0, 93, 332, 1, 0, 0, 0, 95, 334, 1, 0, 0, 0, 97, 336, 1, 0, 0, 0, 99, 338,
		1, 0, 0, 0, 101, 340, 1, 0, 0, 0, 103, 342, 1, 0, 0, 0, 105, 344, 1, 0,
		0, 0, 107, 346, 1, 0, 0, 0, 109, 348, 1, 0, 0, 0, 111, 350, 1, 0, 0, 0,
		113, 352, 1, 0, 0, 0, 115, 354, 1, 0, 0, 0, 117, 356, 1, 0, 0, 0, 119,
		358, 1, 0, 0, 0, 121, 360, 1, 0, 0, 0, 123, 362, 1, 0, 0, 0, 125, 364,
		1, 0, 0, 0, 127, 366, 1, 0, 0, 0, 129, 368, 1, 0, 0, 0, 131, 370, 1, 0,
		0, 0, 133, 372, 1, 0, 0, 0, 135, 374, 1, 0, 0, 0, 137, 376, 1, 0, 0, 0,
		139, 378, 1, 0, 0, 0, 141, 380, 1, 0, 0, 0, 143, 382, 1, 0, 0, 0, 145,
		146, 5, 59, 0, 0, 146, 2, 1, 0, 0, 0, 147, 148, 5, 44, 0, 0, 148, 4, 1,
		0, 0, 0, 149, 150, 5, 40, 0, 0, 150, 6, 1, 0, 0, 0, 151, 152, 5, 41, 0,
		0, 152, 8, 1, 0, 0, 0, 153, 154, 5, 45, 0, 0, 154, 10, 1, 0, 0, 0, 155,
		156, 5, 43, 0, 0, 156, 12, 1, 0, 0, 0, 157, 158, 5, 126, 0, 0, 158, 14,
		1, 0, 0, 0, 159, 160, 5, 124, 0, 0, 160, 161, 5, 124, 0, 0, 161, 16, 1,
		0, 0, 0, 162, 163, 5, 42, 0, 0, 163, 18, 1, 0, 0, 0, 164, 165, 5, 47, 0,
		0, 165, 20, 1, 0, 0, 0, 166, 167, 5, 37, 0, 0, 167, 22, 1, 0, 0, 0, 168,
		169, 5, 60, 0, 0, 169, 170, 5, 60, 0, 0, 170, 24, 1, 0, 0, 0, 171, 172,
		5, 62, 0, 0, 172, 173, 5, 62, 0, 0, 173, 26, 1, 0, 0, 0, 174, 175, 5, 38,
		0, 0, 175, 28, 1, 0, 0, 0, 176, 177, 5, 124, 0, 0, 177, 30, 1, 0, 0, 0,
		178, 179, 5, 60, 0, 0, 179, 32, 1, 0, 0, 0, 180, 181, 5, 60, 0, 0, 181,
		182, 5, 61, 0, 0, 182, 34, 1, 0, 0, 0, 183, 184, 5, 62, 0, 0, 184, 36,
		1, 0, 0, 0, 185, 186, 5, 62, 0, 0, 186, 187, 5, 61, 0, 0, 187, 38, 1, 0,
		0, 0, 188, 189, 5, 61, 0, 0, 189, 40, 1, 0, 0, 0, 190, 191, 5, 61, 0, 0,
		191, 192, 5, 61, 0, 0, 192, 42, 1, 0, 0, 0, 193, 194, 5, 33, 0, 0, 194,
		195, 5, 61, 0, 0, 195, 44, 1, 0, 0, 0, 196, 197, 5, 60, 0, 0, 197, 198,
		5, 62, 0, 0, 198, 46, 1, 0, 0, 0, 199, 200, 5, 36, 0, 0, 200, 201, 5, 107,
		0, 0, 201, 202, 5, 101, 0, 0, 202, 203, 5, 121, 0, 0, 203, 48, 1, 0, 0,
		0, 204, 205, 5, 36, 0, 0, 205, 206, 5, 118, 0, 0, 206, 207, 5, 97, 0, 0,
		207, 208, 5, 108, 0, 0, 208, 209, 5, 117, 0, 0, 209, 210, 5, 101, 0, 0,
		210, 50, 1, 0, 0, 0, 211, 212, 3, 129, 64, 0, 212, 213, 3, 101, 50, 0,
		213, 214, 3, 115, 57, 0, 214, 215, 3, 101, 50, 0, 215, 216, 3, 97, 48,
		0, 216, 217, 3, 131, 65, 0, 217, 52, 1, 0, 0, 0, 218, 219, 3, 103, 51,
		0, 219, 220, 3, 127, 63, 0, 220, 221, 3, 121, 60, 0, 221, 222, 3, 117,
		58, 0, 222, 54, 1, 0, 0, 0, 223, 224, 3, 137, 68, 0, 224, 225, 3, 107,
		53, 0, 225, 226, 3, 101, 50, 0, 226, 227, 3, 127, 63, 0, 227, 228, 3, 101,
		50, 0, 228, 56, 1, 0, 0, 0, 229, 230, 3, 121, 60, 0, 230, 231, 3, 127,
		63, 0, 231, 232, 3, 99, 49, 0, 232, 233, 3, 101, 50, 0, 233, 234, 3, 127,
		63, 0, 234, 58, 1, 0, 0, 0, 235, 236, 3, 95, 47, 0, 236, 237, 3, 141, 70,
		0, 237, 60, 1, 0, 0, 0, 238, 239, 3, 115, 57, 0, 239, 240, 3, 109, 54,
		0, 240, 241, 3, 117, 58, 0, 241, 242, 3, 109, 54, 0, 242, 243, 3, 131,
		65, 0, 243, 62, 1, 0, 0, 0, 244, 245, 3, 93, 46, 0, 245, 246, 3, 129, 64,
		0, 246, 247, 3, 97, 48, 0, 247, 64, 1, 0, 0, 0, 248, 249, 3, 99, 49, 0,
		249, 250, 3, 101, 50, 0, 250, 251, 3, 129, 64, 0, 251, 252, 3, 97, 48,
		0, 252, 66, 1, 0, 0, 0, 253, 254, 3, 109, 54, 0, 254, 255, 3, 129, 64,
		0, 255, 68, 1, 0, 0, 0, 256, 257, 3, 119, 59, 0, 257, 258, 3, 121, 60,
		0, 258, 259, 3, 131, 65, 0, 259, 70, 1, 0, 0, 0, 260, 261, 3, 119, 59,
		0, 261, 262, 3, 133, 66, 0, 262, 263, 3, 115, 57, 0, 263, 264, 3, 115,
		57, 0, 264, 72, 1, 0, 0, 0, 265, 266, 3, 115, 57, 0, 266, 267, 3, 109,
		54, 0, 267, 268, 3, 113, 56, 0, 268, 269, 3, 101, 50, 0, 269, 74, 1, 0,
		0, 0, 270, 271, 3, 93, 46, 0, 271, 272, 3, 119, 59, 0, 272, 273, 3, 99,
		49, 0, 273, 76, 1, 0, 0, 0, 274, 275, 3, 121, 60, 0, 275, 276, 3, 127,
		63, 0, 276, 78, 1, 0, 0, 0, 277, 278, 3, 49, 24, 0, 278, 279, 5, 46, 0,
		0, 279, 80, 1, 0, 0, 0, 280, 285, 7, 0, 0, 0, 281, 284, 3, 83, 41, 0, 282,
		284, 7, 1, 0, 0, 283, 281, 1, 0, 0, 0, 283, 282, 1, 0, 0, 0, 284, 287,
		1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 82, 1, 0,
		0, 0, 287, 285, 1, 0, 0, 0, 288, 289, 7, 2, 0, 0, 289, 84, 1, 0, 0, 0,
		290, 292, 3, 91, 45, 0, 291, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293,
		291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 302, 1, 0, 0, 0, 295, 299,
		5, 46, 0, 0, 296, 298, 3, 91, 45, 0, 297, 296, 1, 0, 0, 0, 298, 301, 1,
		0, 0, 0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 303, 1, 0, 0,
		0, 301, 299, 1, 0, 0, 0, 302, 295, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303,
		311, 1, 0, 0, 0, 304, 306, 5, 46, 0, 0, 305, 307, 3, 91, 45, 0, 306, 305,
		1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0,
		0, 0, 309, 311, 1, 0, 0, 0, 310, 291, 1, 0, 0, 0, 310, 304, 1, 0, 0, 0,
		311, 86, 1, 0, 0, 0, 312, 318, 5, 39, 0, 0, 313, 317, 8, 3, 0, 0, 314,
		315, 5, 39, 0, 0, 315, 317, 5, 39, 0, 0, 316, 313, 1, 0, 0, 0, 316, 314,
		1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0,
		0, 0, 319, 321, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 321, 322, 5, 39, 0, 0,
		322, 88, 1, 0, 0, 0, 323, 325, 7, 4, 0, 0, 324, 323, 1, 0, 0, 0, 325, 326,
		1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 1, 0,
		0, 0, 328, 329, 6, 44, 0, 0, 329, 90, 1, 0, 0, 0, 330, 331, 7, 5, 0, 0,
		331, 92, 1, 0, 0, 0, 332, 333, 7, 6, 0, 0, 333, 94, 1, 0, 0, 0, 334, 335,
		7, 7, 0, 0, 335, 96, 1, 0, 0, 0, 336, 337, 7, 8, 0, 0, 337, 98, 1, 0, 0,
		0, 338, 339, 7, 9, 0, 0, 339, 100, 1, 0, 0, 0, 340, 341, 7, 10, 0, 0, 341,
		102, 1, 0, 0, 0, 342, 343, 7, 11, 0, 0, 343, 104, 1, 0, 0, 0, 344, 345,
		7, 12, 0, 0, 345, 106, 1, 0, 0, 0, 346, 347, 7, 13, 0, 0, 347, 108, 1,
		0, 0, 0, 348, 349, 7, 14, 0, 0, 349, 110, 1, 0, 0, 0, 350, 351, 7, 15,
		0, 0, 351, 112, 1, 0, 0, 0, 352, 353, 7, 16, 0, 0, 353, 114, 1, 0, 0, 0,
		354, 355, 7, 17, 0, 0, 355, 116, 1, 0, 0, 0, 356, 357, 7, 18, 0, 0, 357,
		118, 1, 0, 0, 0, 358, 359, 7, 19, 0, 0, 359, 120, 1, 0, 0, 0, 360, 361,
		7, 20, 0, 0, 361, 122, 1, 0, 0, 0, 362, 363, 7, 21, 0, 0, 363, 124, 1,
		0, 0, 0, 364, 365, 7, 22, 0, 0, 365, 126, 1, 0, 0, 0, 366, 367, 7, 23,
		0, 0, 367, 128, 1, 0, 0, 0, 368, 369, 7, 24, 0, 0, 369, 130, 1, 0, 0, 0,
		370, 371, 7, 25, 0, 0, 371, 132, 1, 0, 0, 0, 372, 373, 7, 26, 0, 0, 373,
		134, 1, 0, 0, 0, 374, 375, 7, 27, 0, 0, 375, 136, 1, 0, 0, 0, 376, 377,
		7, 28, 0, 0, 377, 138, 1, 0, 0, 0, 378, 379, 7, 29, 0, 0, 379, 140, 1,
		0, 0, 0, 380, 381, 7, 30, 0, 0, 381, 142, 1, 0, 0, 0, 382, 383, 7, 31,
		0, 0, 383, 144, 1, 0, 0, 0, 11, 0, 283, 285, 293, 299, 302, 308, 310, 316,
		318, 326, 1, 6, 0, 0,
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

// DSQLLexerInit initializes any static state used to implement DSQLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDSQLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DSQLLexerInit() {
	staticData := &DSQLLexerLexerStaticData
	staticData.once.Do(dsqllexerLexerInit)
}

// NewDSQLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDSQLLexer(input antlr.CharStream) *DSQLLexer {
	DSQLLexerInit()
	l := new(DSQLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DSQLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "DSQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DSQLLexer tokens.
const (
	DSQLLexerT__0                     = 1
	DSQLLexerT__1                     = 2
	DSQLLexerT__2                     = 3
	DSQLLexerT__3                     = 4
	DSQLLexerT__4                     = 5
	DSQLLexerT__5                     = 6
	DSQLLexerT__6                     = 7
	DSQLLexerT__7                     = 8
	DSQLLexerT__8                     = 9
	DSQLLexerT__9                     = 10
	DSQLLexerT__10                    = 11
	DSQLLexerT__11                    = 12
	DSQLLexerT__12                    = 13
	DSQLLexerT__13                    = 14
	DSQLLexerT__14                    = 15
	DSQLLexerT__15                    = 16
	DSQLLexerT__16                    = 17
	DSQLLexerT__17                    = 18
	DSQLLexerT__18                    = 19
	DSQLLexerT__19                    = 20
	DSQLLexerT__20                    = 21
	DSQLLexerT__21                    = 22
	DSQLLexerT__22                    = 23
	DSQLLexerKEY_TOKEN                = 24
	DSQLLexerVALUE_TOKEN              = 25
	DSQLLexerK_SELECT                 = 26
	DSQLLexerK_FROM                   = 27
	DSQLLexerK_WHERE                  = 28
	DSQLLexerK_ORDER                  = 29
	DSQLLexerK_BY                     = 30
	DSQLLexerK_LIMIT                  = 31
	DSQLLexerK_ASC                    = 32
	DSQLLexerK_DESC                   = 33
	DSQLLexerK_IS                     = 34
	DSQLLexerK_NOT                    = 35
	DSQLLexerK_NULL                   = 36
	DSQLLexerK_LIKE                   = 37
	DSQLLexerK_AND                    = 38
	DSQLLexerK_OR                     = 39
	DSQLLexerVALUE_TOKEN_TRAILING_DOT = 40
	DSQLLexerIDENTIFIER               = 41
	DSQLLexerNUMERIC_LITERAL          = 42
	DSQLLexerSTRING_LITERAL           = 43
	DSQLLexerWHITESPACE               = 44
)