package main

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
	default:

	}
}

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
