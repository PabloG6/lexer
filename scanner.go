package main

type Scanner struct {
	Source  string
	Tokens  []Token
	current int
	start   int
	line    int
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
		if s.current >= len(s.Tokens) {
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
	case '*':
		s.AddToken(STAR)
		break
	}
}

func (s *Scanner) advanceCharacter() uint8 {
	s.current++
	return s.Source[s.current]
}

func (s *Scanner) AddToken(t TOKEN) {
	text := s.Source[s.start:s.current]
}
