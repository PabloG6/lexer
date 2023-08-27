package main

type TOKEN int64

const (
	LEFT_PAREN TOKEN = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	SEMI_COLON
	SLASH
	STAR
	MINUS
	PLUS
	SEMICOLON
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	IDENTIFIER
	STRING
	NUMBER

	AND
	CLASS
	ELSE
	IF
	IF_ELSE
	TRUE
	FALSE
	FUN
	OR
	PRINT
	RETURN
	SUPER
	THIS
	VAR
	WHILE
	EOF
)

type TokenType struct {
	Type    TOKEN
	Lexeme  string
	Literal interface{}
	Line    int64
}

func NewTokenType(token TOKEN, lexeme string, line int64) TokenType {
	return TokenType{
		Type:   token,
		Lexeme: lexeme,
		Line:   line,
	}
}
