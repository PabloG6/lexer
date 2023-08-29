package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	Source  string
	Tokens  []TokenType
	current int
	start   int
	line    int64
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		current: 0,
		start:   0,
		line:    1,
	}
}
func (s *Lexer) ScanTokens() {
	for {
		if s.IsAtEnd() {
			break
		}

		s.ScanToken()
	}
}

func (s *Lexer) ScanToken() {
	var c = s.advanceCharacter()
	switch c {
	case '(':
		s.AddToken(LEFT_PAREN)
		break
	case ')':
		s.AddToken(RIGHT_PAREN)
		break
	case '{':
		s.AddToken(LEFT_BRACE)
		break
	case '}':
		s.AddToken(RIGHT_BRACE)
		break
	case ',':
		s.AddToken(COMMA)
		break
	case '.':
		s.AddToken(DOT)
		break
	case '-':
		s.AddToken(MINUS)
		break
	case '+':
		s.AddToken(PLUS)
		break
	case ';':
		s.AddToken(SEMICOLON)
		break
	case '!':
		if s.match('=') {
			s.AddToken(BANG_EQUAL)
		} else {
			s.AddToken(BANG)
		}
		break

	case '=':
		break
	case '>':
		if s.match('=') {
			s.AddToken(GREATER_EQUAL)
		} else {
			s.AddToken(GREATER)
		}
		break
	case '/':
		if s.match('/') {
			for {
				if s.peek() != '\n' && !s.IsAtEnd() {
					s.advanceCharacter()
					continue
				}
				break
			}
		} else {
			s.AddToken(SLASH)
		}

		break
	case '<':
		if s.match('=') {
			s.AddToken(LESS_EQUAL)
		} else {
			s.AddToken(LESS)
		}

		break
	case '*':
		s.AddToken(STAR)
		break
	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace.
		break
	case '\n':
		s.line++
		break

	case '"':
		for {
			if s.peek() != '"' && !s.IsAtEnd() {
				if s.peek() != '\n' {
					s.line++
				}

				s.advanceCharacter()

			} else {
				break
			}

		}

		if s.IsAtEnd() {
			ReportError(s.line, "String without closing parenthesis")
		}

		val := s.Source[s.start+1 : s.current+1]
		s.AddStringToken(STRING, val)
	default:
		if s.IsNumber(c) {
			s.CaptureNumber()
		} else if s.IsAlpha(c) {

		} else {
			ReportError(s.line, fmt.Sprintf("Unexpected character: %c", c))
		}

	}
}

/*
*
look ahead . to see if characters are numbers.
*/
func (s *Lexer) peekNext() rune {
	if s.current+1 >= len(s.Source) {
		return rune(0)

	}

	return rune(s.Source[s.current+1])
}

func (s *Lexer) CaptureNumber() {
	for {
		if s.IsNumber(s.peek()) {
			s.advanceCharacter()
			if s.peek() == '.' && s.IsNumber(s.peekNext()) {
				s.advanceCharacter()

				for {
					if s.IsNumber(s.peek()) {
						s.advanceCharacter()
					}
				}
			}
		}
	}
}
func (s *Lexer) IsNumber(c rune) bool {

	return unicode.IsDigit(c)
}

func (s *Lexer) AddStringToken(t TOKEN, val string) {
	s.Tokens = append(s.Tokens, NewTokenType(t, val, s.line))
}

func (s *Lexer) advanceCharacter() rune {
	s.current++
	return rune(s.Source[s.current])
}

func (s *Lexer) AddToken(t TOKEN) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, NewTokenType(t, text, s.line))
}

func (s *Lexer) IsAtEnd() bool {
	return s.current >= len(s.Source)
}
func (s *Lexer) match(c rune) bool {
	if s.IsAtEnd() {
		return false
	}

	if rune(s.Source[s.current]) != c {
		return false
	}

	s.current++
	return true
}

func (s *Lexer) peek() rune {
	if s.IsAtEnd() {
		return rune(0)
	}

	return rune(s.Source[s.current])
}

func (s *Lexer) IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (s *Lexer) isAlphaNumeric(c rune) bool {
	return s.IsAlpha(c) || s.IsNumber(c)
}
