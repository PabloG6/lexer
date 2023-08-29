package main

import "lexer/lexer"

var RESERVE_WORDS = make(map[string]lexer.TOKEN)

func init() {
	RESERVE_WORDS["else"] = lexer.ELSE
	RESERVE_WORDS["and"] = lexer.AND
	RESERVE_WORDS["class"] = lexer.CLASS
	RESERVE_WORDS["else"] = lexer.ELSE
	RESERVE_WORDS["false"] = lexer.FALSE
	RESERVE_WORDS["for"] = lexer.FOR
	RESERVE_WORDS["fun"] = lexer.FUN
	RESERVE_WORDS["if"] = lexer.IF
	RESERVE_WORDS["nil"] = lexer.NIL
	RESERVE_WORDS["or"] = lexer.OR
	RESERVE_WORDS["print"] = lexer.PRINT
	RESERVE_WORDS["return"] = lexer.RETURN
	RESERVE_WORDS["super"] = lexer.SUPER
	RESERVE_WORDS["this"] = lexer.THIS
	RESERVE_WORDS["true"] = lexer.TRUE
	RESERVE_WORDS["var"] = lexer.VAR
	RESERVE_WORDS["while"] = lexer.WHILE
}
func (s *lexer.Scanner) parseIdentifier() {
	for {
		if s.isAlphaNumeric(s.peek()) {
			s.advanceCharacter()
			continue
		}

		keyword := s.Source[s.start:s.current]
		identifier, ok := RESERVE_WORDS[keyword]
		if ok {
			s.AddToken(identifier)
		} else {
			s.AddToken(lexer.IDENTIFIER)
		}

		break
	}
}
