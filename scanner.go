package main

<<<<<<< HEAD
=======
import (
	"fmt"
	"unicode"
)

>>>>>>> a8c20af (add alphanumeric functions for reserved keywords)
type Scanner struct {
	Source  string
	Tokens  []TokenType
	current int
	start   int
	line    int64
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		current: 0,
		start:   0,
		line:    1,
	}
}
func (s *Scanner) ScanTokens() {
	for {
		if s.IsAtEnd() {
			break
		}

		s.ScanToken()
	}
}

func (s *Scanner) ScanToken() {
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
<<<<<<< HEAD
	default:
=======
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
>>>>>>> a8c20af (add alphanumeric functions for reserved keywords)

	}
}

<<<<<<< HEAD
=======
/*
*
look ahead . characters is numbers.
*/
func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.Source) {
		return rune(0)

	}

	return rune(s.Source[s.current+1])
}

func (s *Scanner) CaptureNumber() {
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
func (s *Scanner) IsNumber(c rune) bool {

	return unicode.IsDigit(c)
}

func (s *Scanner) AddStringToken(t TOKEN, val string) {
	s.Tokens = append(s.Tokens, NewTokenType(t, val, s.line))
}

>>>>>>> a8c20af (add alphanumeric functions for reserved keywords)
func (s *Scanner) advanceCharacter() rune {
	s.current++
	return rune(s.Source[s.current])
}

func (s *Scanner) AddToken(t TOKEN) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, NewTokenType(t, text, s.line))
}

func (s *Scanner) IsAtEnd() bool {
	return s.current >= len(s.Source)
}
func (s *Scanner) match(c rune) bool {
	if s.IsAtEnd() {
		return false
	}

	if rune(s.Source[s.current]) != c {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() rune {
	if s.IsAtEnd() {
		return rune(0)
	}

	return rune(s.Source[s.current])
}
<<<<<<< HEAD
=======

func (s *Scanner) IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (s *Scanner) isAlphaNumeric(c rune) bool {
	return s.IsAlpha(c) || s.IsNumber(c)
}
>>>>>>> a8c20af (add alphanumeric functions for reserved keywords)
