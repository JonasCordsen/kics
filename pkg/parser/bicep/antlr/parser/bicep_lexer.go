// Code generated from bicep.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
	"bufio"
	"strings"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type bicepLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

type BlockInfo struct {
    BlockType string
    StartLine int
    EndLine   int
}

type LineInfo struct {
    Type   string
    Bytes  struct {
        Bytes  []byte
        String string
    }
    Range struct {
        Start struct {
            Line   int
            Column int
            Byte   int
        }
        End struct {
            Line   int
            Column int
            Byte   int
        }
    }
	Block BlockInfo
}

func (l *bicepLexer) GetLinesInfo() []LineInfo {
    var linesInfo []LineInfo
    input := l.GetInputStream().GetText(0, l.GetInputStream().Size()-1)
    scanner := bufio.NewScanner(strings.NewReader(input))
    lineNumber := 0
    byteOffset := 0

    var currentBlock BlockInfo
    blockStack := []BlockInfo{}

    for scanner.Scan() {
        line := scanner.Text()
        trimmedLine := strings.TrimSpace(line) // Trim leading and trailing whitespace
        lineBytes := []byte(line)
        lineLength := len(lineBytes)

        // Create a new lexer instance for the current line
        lineLexer := NewbicepLexer(antlr.NewInputStream(line))
        tokens := lineLexer.GetAllTokens()

		if len(blockStack) > 0 {
            currentBlock = blockStack[len(blockStack)-1]
        }

		if strings.HasSuffix(trimmedLine, "{") {
			currentBlock = BlockInfo{
                BlockType: trimmedLine,
                StartLine: lineNumber,
            }
            blockStack = append(blockStack, currentBlock)
		} else if strings.Contains(trimmedLine, "}") {
			if len(blockStack) > 0 {
				currentBlock.EndLine = lineNumber
				for i := range linesInfo {
					if linesInfo[i].Block.StartLine == currentBlock.StartLine {
						linesInfo[i].Block.EndLine = lineNumber
					}
				}
				blockStack = blockStack[:len(blockStack)-1]
			}
		}

        // Determine the type of the line based on the tokens
        lineType := "other"
        if len(tokens) > 0 {
            lineType = l.SymbolicNames[tokens[0].GetTokenType()]
        }

        lineInfo := LineInfo{
            Type: lineType,
            Bytes: struct {
                Bytes  []byte
                String string
            }{
                Bytes:  lineBytes,
                String: line,
            },
            Range: struct {
                Start struct {
                    Line   int
                    Column int
                    Byte   int
                }
                End struct {
                    Line   int
                    Column int
                    Byte   int
                }
            }{
                Start: struct {
                    Line   int
                    Column int
                    Byte   int
                }{
                    Line:   lineNumber,
                    Column: 0,
                    Byte:   byteOffset,
                },
                End: struct {
                    Line   int
                    Column int
                    Byte   int
                }{
                    Line:   lineNumber,
                    Column: lineLength,
                    Byte:   byteOffset + lineLength,
                },
            },
            Block: currentBlock,
        }

        linesInfo = append(linesInfo, lineInfo)
        lineNumber++
        byteOffset += lineLength + 1 // +1 for the newline character
    }

    return linesInfo
}

var BicepLexerLexerStaticData struct {
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

func biceplexerLexerInit() {
	staticData := &BicepLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'@'", "','", "'['", "']'", "'('", "')'", "'.'", "'|'", "",
		"'='", "'{'", "'}'", "'param'", "'var'", "'true'", "'false'", "'null'",
		"'array'", "'object'", "'resource'", "'output'", "'targetScope'", "'import'",
		"'with'", "'as'", "'metadata'", "'existing'", "'type'", "'module'",
		"", "", "", "", "'string'", "'int'", "'bool'", "'if'", "'for'", "'in'",
		"'?'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='", "'=>'",
	}
	staticData.SymbolicNames = []string{
		"", "MULTILINE_STRING", "AT", "COMMA", "OBRACK", "CBRACK", "OPAR", "CPAR",
		"DOT", "PIPE", "COL", "ASSIGN", "OBRACE", "CBRACE", "PARAM", "VAR",
		"TRUE", "FALSE", "NULL", "ARRAY", "OBJECT", "RESOURCE", "OUTPUT", "TARGET_SCOPE",
		"IMPORT", "WITH", "AS", "METADATA", "EXISTING", "TYPE", "MODULE", "STRING_LEFT_PIECE",
		"STRING_MIDDLE_PIECE", "STRING_RIGHT_PIECE", "STRING_COMPLETE", "STRING",
		"INT", "BOOL", "IF", "FOR", "IN", "QMARK", "GT", "GTE", "LT", "LTE",
		"EQ", "NEQ", "ARROW", "IDENTIFIER", "NUMBER", "NL", "SINGLE_LINE_COMMENT",
		"MULTI_LINE_COMMENT", "SPACES", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"MULTILINE_STRING", "AT", "COMMA", "OBRACK", "CBRACK", "OPAR", "CPAR",
		"DOT", "PIPE", "COL", "ASSIGN", "OBRACE", "CBRACE", "PARAM", "VAR",
		"TRUE", "FALSE", "NULL", "ARRAY", "OBJECT", "RESOURCE", "OUTPUT", "TARGET_SCOPE",
		"IMPORT", "WITH", "AS", "METADATA", "EXISTING", "TYPE", "MODULE", "STRING_LEFT_PIECE",
		"STRING_MIDDLE_PIECE", "STRING_RIGHT_PIECE", "STRING_COMPLETE", "STRING",
		"INT", "BOOL", "IF", "FOR", "IN", "QMARK", "GT", "GTE", "LT", "LTE",
		"EQ", "NEQ", "ARROW", "IDENTIFIER", "NUMBER", "NL", "SINGLE_LINE_COMMENT",
		"MULTI_LINE_COMMENT", "SPACES", "UNKNOWN", "STRINGCHAR", "ESCAPE", "HEX",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 55, 434, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 57, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 123, 8, 0, 10, 0, 12, 0, 126,
		9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		3, 9, 151, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 273, 8, 30, 10,
		30, 12, 30, 276, 9, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 283,
		8, 31, 10, 31, 12, 31, 286, 9, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5,
		32, 293, 8, 32, 10, 32, 12, 32, 296, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		5, 33, 302, 8, 33, 10, 33, 12, 33, 305, 9, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 5, 48, 358, 8, 48, 10, 48, 12, 48,
		361, 9, 48, 1, 49, 4, 49, 364, 8, 49, 11, 49, 12, 49, 365, 1, 49, 1, 49,
		4, 49, 370, 8, 49, 11, 49, 12, 49, 371, 3, 49, 374, 8, 49, 1, 50, 4, 50,
		377, 8, 50, 11, 50, 12, 50, 378, 1, 51, 1, 51, 1, 51, 1, 51, 5, 51, 385,
		8, 51, 10, 51, 12, 51, 388, 9, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1,
		52, 5, 52, 396, 8, 52, 10, 52, 12, 52, 399, 9, 52, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 53, 4, 53, 407, 8, 53, 11, 53, 12, 53, 408, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 55, 1, 55, 3, 55, 417, 8, 55, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 4, 56, 425, 8, 56, 11, 56, 12, 56, 426, 1, 56, 1,
		56, 3, 56, 431, 8, 56, 1, 57, 1, 57, 2, 124, 397, 0, 58, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97,
		49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 0, 113, 0,
		115, 0, 1, 0, 8, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 5,
		0, 9, 10, 13, 13, 36, 36, 39, 39, 92, 92, 6, 0, 36, 36, 39, 39, 92, 92,
		110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 447, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1,
		0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101,
		1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0,
		0, 109, 1, 0, 0, 0, 1, 117, 1, 0, 0, 0, 3, 131, 1, 0, 0, 0, 5, 133, 1,
		0, 0, 0, 7, 135, 1, 0, 0, 0, 9, 137, 1, 0, 0, 0, 11, 139, 1, 0, 0, 0, 13,
		141, 1, 0, 0, 0, 15, 143, 1, 0, 0, 0, 17, 145, 1, 0, 0, 0, 19, 150, 1,
		0, 0, 0, 21, 152, 1, 0, 0, 0, 23, 154, 1, 0, 0, 0, 25, 156, 1, 0, 0, 0,
		27, 158, 1, 0, 0, 0, 29, 164, 1, 0, 0, 0, 31, 168, 1, 0, 0, 0, 33, 173,
		1, 0, 0, 0, 35, 179, 1, 0, 0, 0, 37, 184, 1, 0, 0, 0, 39, 190, 1, 0, 0,
		0, 41, 197, 1, 0, 0, 0, 43, 206, 1, 0, 0, 0, 45, 213, 1, 0, 0, 0, 47, 225,
		1, 0, 0, 0, 49, 232, 1, 0, 0, 0, 51, 237, 1, 0, 0, 0, 53, 240, 1, 0, 0,
		0, 55, 249, 1, 0, 0, 0, 57, 258, 1, 0, 0, 0, 59, 263, 1, 0, 0, 0, 61, 270,
		1, 0, 0, 0, 63, 280, 1, 0, 0, 0, 65, 290, 1, 0, 0, 0, 67, 299, 1, 0, 0,
		0, 69, 308, 1, 0, 0, 0, 71, 315, 1, 0, 0, 0, 73, 319, 1, 0, 0, 0, 75, 324,
		1, 0, 0, 0, 77, 327, 1, 0, 0, 0, 79, 331, 1, 0, 0, 0, 81, 334, 1, 0, 0,
		0, 83, 336, 1, 0, 0, 0, 85, 338, 1, 0, 0, 0, 87, 341, 1, 0, 0, 0, 89, 343,
		1, 0, 0, 0, 91, 346, 1, 0, 0, 0, 93, 349, 1, 0, 0, 0, 95, 352, 1, 0, 0,
		0, 97, 355, 1, 0, 0, 0, 99, 363, 1, 0, 0, 0, 101, 376, 1, 0, 0, 0, 103,
		380, 1, 0, 0, 0, 105, 391, 1, 0, 0, 0, 107, 406, 1, 0, 0, 0, 109, 412,
		1, 0, 0, 0, 111, 416, 1, 0, 0, 0, 113, 418, 1, 0, 0, 0, 115, 432, 1, 0,
		0, 0, 117, 118, 5, 39, 0, 0, 118, 119, 5, 39, 0, 0, 119, 120, 5, 39, 0,
		0, 120, 124, 1, 0, 0, 0, 121, 123, 9, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123,
		126, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 127,
		1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 5, 39, 0, 0, 128, 129, 5, 39,
		0, 0, 129, 130, 5, 39, 0, 0, 130, 2, 1, 0, 0, 0, 131, 132, 5, 64, 0, 0,
		132, 4, 1, 0, 0, 0, 133, 134, 5, 44, 0, 0, 134, 6, 1, 0, 0, 0, 135, 136,
		5, 91, 0, 0, 136, 8, 1, 0, 0, 0, 137, 138, 5, 93, 0, 0, 138, 10, 1, 0,
		0, 0, 139, 140, 5, 40, 0, 0, 140, 12, 1, 0, 0, 0, 141, 142, 5, 41, 0, 0,
		142, 14, 1, 0, 0, 0, 143, 144, 5, 46, 0, 0, 144, 16, 1, 0, 0, 0, 145, 146,
		5, 124, 0, 0, 146, 18, 1, 0, 0, 0, 147, 151, 5, 58, 0, 0, 148, 149, 5,
		58, 0, 0, 149, 151, 5, 58, 0, 0, 150, 147, 1, 0, 0, 0, 150, 148, 1, 0,
		0, 0, 151, 20, 1, 0, 0, 0, 152, 153, 5, 61, 0, 0, 153, 22, 1, 0, 0, 0,
		154, 155, 5, 123, 0, 0, 155, 24, 1, 0, 0, 0, 156, 157, 5, 125, 0, 0, 157,
		26, 1, 0, 0, 0, 158, 159, 5, 112, 0, 0, 159, 160, 5, 97, 0, 0, 160, 161,
		5, 114, 0, 0, 161, 162, 5, 97, 0, 0, 162, 163, 5, 109, 0, 0, 163, 28, 1,
		0, 0, 0, 164, 165, 5, 118, 0, 0, 165, 166, 5, 97, 0, 0, 166, 167, 5, 114,
		0, 0, 167, 30, 1, 0, 0, 0, 168, 169, 5, 116, 0, 0, 169, 170, 5, 114, 0,
		0, 170, 171, 5, 117, 0, 0, 171, 172, 5, 101, 0, 0, 172, 32, 1, 0, 0, 0,
		173, 174, 5, 102, 0, 0, 174, 175, 5, 97, 0, 0, 175, 176, 5, 108, 0, 0,
		176, 177, 5, 115, 0, 0, 177, 178, 5, 101, 0, 0, 178, 34, 1, 0, 0, 0, 179,
		180, 5, 110, 0, 0, 180, 181, 5, 117, 0, 0, 181, 182, 5, 108, 0, 0, 182,
		183, 5, 108, 0, 0, 183, 36, 1, 0, 0, 0, 184, 185, 5, 97, 0, 0, 185, 186,
		5, 114, 0, 0, 186, 187, 5, 114, 0, 0, 187, 188, 5, 97, 0, 0, 188, 189,
		5, 121, 0, 0, 189, 38, 1, 0, 0, 0, 190, 191, 5, 111, 0, 0, 191, 192, 5,
		98, 0, 0, 192, 193, 5, 106, 0, 0, 193, 194, 5, 101, 0, 0, 194, 195, 5,
		99, 0, 0, 195, 196, 5, 116, 0, 0, 196, 40, 1, 0, 0, 0, 197, 198, 5, 114,
		0, 0, 198, 199, 5, 101, 0, 0, 199, 200, 5, 115, 0, 0, 200, 201, 5, 111,
		0, 0, 201, 202, 5, 117, 0, 0, 202, 203, 5, 114, 0, 0, 203, 204, 5, 99,
		0, 0, 204, 205, 5, 101, 0, 0, 205, 42, 1, 0, 0, 0, 206, 207, 5, 111, 0,
		0, 207, 208, 5, 117, 0, 0, 208, 209, 5, 116, 0, 0, 209, 210, 5, 112, 0,
		0, 210, 211, 5, 117, 0, 0, 211, 212, 5, 116, 0, 0, 212, 44, 1, 0, 0, 0,
		213, 214, 5, 116, 0, 0, 214, 215, 5, 97, 0, 0, 215, 216, 5, 114, 0, 0,
		216, 217, 5, 103, 0, 0, 217, 218, 5, 101, 0, 0, 218, 219, 5, 116, 0, 0,
		219, 220, 5, 83, 0, 0, 220, 221, 5, 99, 0, 0, 221, 222, 5, 111, 0, 0, 222,
		223, 5, 112, 0, 0, 223, 224, 5, 101, 0, 0, 224, 46, 1, 0, 0, 0, 225, 226,
		5, 105, 0, 0, 226, 227, 5, 109, 0, 0, 227, 228, 5, 112, 0, 0, 228, 229,
		5, 111, 0, 0, 229, 230, 5, 114, 0, 0, 230, 231, 5, 116, 0, 0, 231, 48,
		1, 0, 0, 0, 232, 233, 5, 119, 0, 0, 233, 234, 5, 105, 0, 0, 234, 235, 5,
		116, 0, 0, 235, 236, 5, 104, 0, 0, 236, 50, 1, 0, 0, 0, 237, 238, 5, 97,
		0, 0, 238, 239, 5, 115, 0, 0, 239, 52, 1, 0, 0, 0, 240, 241, 5, 109, 0,
		0, 241, 242, 5, 101, 0, 0, 242, 243, 5, 116, 0, 0, 243, 244, 5, 97, 0,
		0, 244, 245, 5, 100, 0, 0, 245, 246, 5, 97, 0, 0, 246, 247, 5, 116, 0,
		0, 247, 248, 5, 97, 0, 0, 248, 54, 1, 0, 0, 0, 249, 250, 5, 101, 0, 0,
		250, 251, 5, 120, 0, 0, 251, 252, 5, 105, 0, 0, 252, 253, 5, 115, 0, 0,
		253, 254, 5, 116, 0, 0, 254, 255, 5, 105, 0, 0, 255, 256, 5, 110, 0, 0,
		256, 257, 5, 103, 0, 0, 257, 56, 1, 0, 0, 0, 258, 259, 5, 116, 0, 0, 259,
		260, 5, 121, 0, 0, 260, 261, 5, 112, 0, 0, 261, 262, 5, 101, 0, 0, 262,
		58, 1, 0, 0, 0, 263, 264, 5, 109, 0, 0, 264, 265, 5, 111, 0, 0, 265, 266,
		5, 100, 0, 0, 266, 267, 5, 117, 0, 0, 267, 268, 5, 108, 0, 0, 268, 269,
		5, 101, 0, 0, 269, 60, 1, 0, 0, 0, 270, 274, 5, 39, 0, 0, 271, 273, 3,
		111, 55, 0, 272, 271, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0,
		0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0,
		277, 278, 5, 36, 0, 0, 278, 279, 5, 123, 0, 0, 279, 62, 1, 0, 0, 0, 280,
		284, 5, 125, 0, 0, 281, 283, 3, 111, 55, 0, 282, 281, 1, 0, 0, 0, 283,
		286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 287,
		1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 287, 288, 5, 36, 0, 0, 288, 289, 5, 123,
		0, 0, 289, 64, 1, 0, 0, 0, 290, 294, 5, 125, 0, 0, 291, 293, 3, 111, 55,
		0, 292, 291, 1, 0, 0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294,
		295, 1, 0, 0, 0, 295, 297, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 297, 298,
		5, 39, 0, 0, 298, 66, 1, 0, 0, 0, 299, 303, 5, 39, 0, 0, 300, 302, 3, 111,
		55, 0, 301, 300, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0,
		303, 304, 1, 0, 0, 0, 304, 306, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306,
		307, 5, 39, 0, 0, 307, 68, 1, 0, 0, 0, 308, 309, 5, 115, 0, 0, 309, 310,
		5, 116, 0, 0, 310, 311, 5, 114, 0, 0, 311, 312, 5, 105, 0, 0, 312, 313,
		5, 110, 0, 0, 313, 314, 5, 103, 0, 0, 314, 70, 1, 0, 0, 0, 315, 316, 5,
		105, 0, 0, 316, 317, 5, 110, 0, 0, 317, 318, 5, 116, 0, 0, 318, 72, 1,
		0, 0, 0, 319, 320, 5, 98, 0, 0, 320, 321, 5, 111, 0, 0, 321, 322, 5, 111,
		0, 0, 322, 323, 5, 108, 0, 0, 323, 74, 1, 0, 0, 0, 324, 325, 5, 105, 0,
		0, 325, 326, 5, 102, 0, 0, 326, 76, 1, 0, 0, 0, 327, 328, 5, 102, 0, 0,
		328, 329, 5, 111, 0, 0, 329, 330, 5, 114, 0, 0, 330, 78, 1, 0, 0, 0, 331,
		332, 5, 105, 0, 0, 332, 333, 5, 110, 0, 0, 333, 80, 1, 0, 0, 0, 334, 335,
		5, 63, 0, 0, 335, 82, 1, 0, 0, 0, 336, 337, 5, 62, 0, 0, 337, 84, 1, 0,
		0, 0, 338, 339, 5, 62, 0, 0, 339, 340, 5, 61, 0, 0, 340, 86, 1, 0, 0, 0,
		341, 342, 5, 60, 0, 0, 342, 88, 1, 0, 0, 0, 343, 344, 5, 60, 0, 0, 344,
		345, 5, 61, 0, 0, 345, 90, 1, 0, 0, 0, 346, 347, 5, 61, 0, 0, 347, 348,
		5, 61, 0, 0, 348, 92, 1, 0, 0, 0, 349, 350, 5, 33, 0, 0, 350, 351, 5, 61,
		0, 0, 351, 94, 1, 0, 0, 0, 352, 353, 5, 61, 0, 0, 353, 354, 5, 62, 0, 0,
		354, 96, 1, 0, 0, 0, 355, 359, 7, 0, 0, 0, 356, 358, 7, 1, 0, 0, 357, 356,
		1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0,
		0, 0, 360, 98, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 364, 7, 2, 0, 0,
		363, 362, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 365,
		366, 1, 0, 0, 0, 366, 373, 1, 0, 0, 0, 367, 369, 5, 46, 0, 0, 368, 370,
		7, 2, 0, 0, 369, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 369, 1, 0,
		0, 0, 371, 372, 1, 0, 0, 0, 372, 374, 1, 0, 0, 0, 373, 367, 1, 0, 0, 0,
		373, 374, 1, 0, 0, 0, 374, 100, 1, 0, 0, 0, 375, 377, 7, 3, 0, 0, 376,
		375, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379,
		1, 0, 0, 0, 379, 102, 1, 0, 0, 0, 380, 381, 5, 47, 0, 0, 381, 382, 5, 47,
		0, 0, 382, 386, 1, 0, 0, 0, 383, 385, 8, 3, 0, 0, 384, 383, 1, 0, 0, 0,
		385, 388, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387,
		389, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 389, 390, 6, 51, 0, 0, 390, 104,
		1, 0, 0, 0, 391, 392, 5, 47, 0, 0, 392, 393, 5, 42, 0, 0, 393, 397, 1,
		0, 0, 0, 394, 396, 9, 0, 0, 0, 395, 394, 1, 0, 0, 0, 396, 399, 1, 0, 0,
		0, 397, 398, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 398, 400, 1, 0, 0, 0, 399,
		397, 1, 0, 0, 0, 400, 401, 5, 42, 0, 0, 401, 402, 5, 47, 0, 0, 402, 403,
		1, 0, 0, 0, 403, 404, 6, 52, 1, 0, 404, 106, 1, 0, 0, 0, 405, 407, 7, 4,
		0, 0, 406, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0,
		408, 409, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 411, 6, 53, 1, 0, 411,
		108, 1, 0, 0, 0, 412, 413, 9, 0, 0, 0, 413, 110, 1, 0, 0, 0, 414, 417,
		8, 5, 0, 0, 415, 417, 3, 113, 56, 0, 416, 414, 1, 0, 0, 0, 416, 415, 1,
		0, 0, 0, 417, 112, 1, 0, 0, 0, 418, 430, 5, 92, 0, 0, 419, 431, 7, 6, 0,
		0, 420, 421, 5, 117, 0, 0, 421, 422, 5, 123, 0, 0, 422, 424, 1, 0, 0, 0,
		423, 425, 3, 115, 57, 0, 424, 423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426,
		424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429,
		5, 125, 0, 0, 429, 431, 1, 0, 0, 0, 430, 419, 1, 0, 0, 0, 430, 420, 1,
		0, 0, 0, 431, 114, 1, 0, 0, 0, 432, 433, 7, 7, 0, 0, 433, 116, 1, 0, 0,
		0, 18, 0, 124, 150, 274, 284, 294, 303, 359, 365, 371, 373, 378, 386, 397,
		408, 416, 426, 430, 2, 0, 1, 0, 6, 0, 0,
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

// bicepLexerInit initializes any static state used to implement bicepLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewbicepLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BicepLexerInit() {
	staticData := &BicepLexerLexerStaticData
	staticData.once.Do(biceplexerLexerInit)
}

// NewbicepLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewbicepLexer(input antlr.CharStream) *bicepLexer {
	BicepLexerInit()
	l := new(bicepLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BicepLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "bicep.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// bicepLexer tokens.
const (
	bicepLexerMULTILINE_STRING    = 1
	bicepLexerAT                  = 2
	bicepLexerCOMMA               = 3
	bicepLexerOBRACK              = 4
	bicepLexerCBRACK              = 5
	bicepLexerOPAR                = 6
	bicepLexerCPAR                = 7
	bicepLexerDOT                 = 8
	bicepLexerPIPE                = 9
	bicepLexerCOL                 = 10
	bicepLexerASSIGN              = 11
	bicepLexerOBRACE              = 12
	bicepLexerCBRACE              = 13
	bicepLexerPARAM               = 14
	bicepLexerVAR                 = 15
	bicepLexerTRUE                = 16
	bicepLexerFALSE               = 17
	bicepLexerNULL                = 18
	bicepLexerARRAY               = 19
	bicepLexerOBJECT              = 20
	bicepLexerRESOURCE            = 21
	bicepLexerOUTPUT              = 22
	bicepLexerTARGET_SCOPE        = 23
	bicepLexerIMPORT              = 24
	bicepLexerWITH                = 25
	bicepLexerAS                  = 26
	bicepLexerMETADATA            = 27
	bicepLexerEXISTING            = 28
	bicepLexerTYPE                = 29
	bicepLexerMODULE              = 30
	bicepLexerSTRING_LEFT_PIECE   = 31
	bicepLexerSTRING_MIDDLE_PIECE = 32
	bicepLexerSTRING_RIGHT_PIECE  = 33
	bicepLexerSTRING_COMPLETE     = 34
	bicepLexerSTRING              = 35
	bicepLexerINT                 = 36
	bicepLexerBOOL                = 37
	bicepLexerIF                  = 38
	bicepLexerFOR                 = 39
	bicepLexerIN                  = 40
	bicepLexerQMARK               = 41
	bicepLexerGT                  = 42
	bicepLexerGTE                 = 43
	bicepLexerLT                  = 44
	bicepLexerLTE                 = 45
	bicepLexerEQ                  = 46
	bicepLexerNEQ                 = 47
	bicepLexerARROW               = 48
	bicepLexerIDENTIFIER          = 49
	bicepLexerNUMBER              = 50
	bicepLexerNL                  = 51
	bicepLexerSINGLE_LINE_COMMENT = 52
	bicepLexerMULTI_LINE_COMMENT  = 53
	bicepLexerSPACES              = 54
	bicepLexerUNKNOWN             = 55
)
