package main

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
	tokens []Token
}

func main() {

}
