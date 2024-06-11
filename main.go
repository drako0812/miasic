package main

import (
	"fmt"
	"os"
	"unicode"
)

const SOURCE string = `EXIT 0;`

type TokenT uint

const (
	TK_INVALID TokenT = iota
	TK_KW
	TK_LIT_INT
	TK_SEMICOLON
)

type Token struct {
	Type     TokenT
	FilePath string
	Line     int64
	Column   int64
}

func NewToken(t TokenT, path string, line, column int64) *Token {
	return &Token{
		Type:     t,
		FilePath: path,
		Line:     line,
		Column:   column,
	}
}

type TokenKW struct {
	*Token
	Value string
}

func NewTokenKW(v string, path string, line, column int64) *TokenKW {
	return &TokenKW{
		Token: NewToken(TK_KW, path, line, column),
		Value: v,
	}
}

type TokenLitInt struct {
	*Token
	Value int32
}

func NewTokenLitInt(v int32, path string, line, column int64) *TokenLitInt {
	return &TokenLitInt{
		Token: NewToken(TK_LIT_INT, path, line, column),
		Value: v,
	}
}

type Tokenizer struct {
	tokens []*Token
	pos int // The current offset into src
	line int // The current line inside the src
	col int // The current column inside the src
	src string // The source code being read
}

func NewTokenizerFromString(filename, source string) *Tokenizer {
	return &Tokenizer {
		tokens: []*Token{},
		pos: 0,
		line: 0,
		col: 0,
		src: source,
	}
}

func NewTokenizer(filepath string) *Tokenizer {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file \"%v\": %v\n", filepath, err)
		return nil
	}

	return NewTokenizerFromString(filepath, string(data))
}

func (t *Tokenizer) Advance() error {
	t.pos++
	if t.pos >= len(t.src) {
		return fmt.Errorf("Unexpected end of file")
	}

	t.col++

	if t.src[t.pos] == '\n' {
		t.line++
		t.col = 0
	}

	return nil
}

func (t *Tokenizer) AdvanceMulti(n int) (int, error) {
	var count int = 0
	for i := 0; i < n; i++ {
		if err := t.Advance(); err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (t *Tokenizer) Peek(offset int) (rune, error) {
	if t.pos + offset >= len(t.src) {
		return 0, fmt.Errorf("Unexpected end of file")
	}

	return rune(t.src[t.pos+offset]), nil
}

func (t *Tokenizer) EatWhitespace() error {
	ch, err := t.Peek(0)
	for unicode.IsSpace(ch) && err == nil {
		if err := t.Advance(); err != nil {
			return err
		}
	}
}

func (t *Tokenizer) Tokenize() []*Token {

	return t.tokens
}

func main() {

}
