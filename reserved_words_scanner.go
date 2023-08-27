package main

var RESERVE_WORDS = make(map[string]TOKEN)

func init() {
	RESERVE_WORDS["else"] = ELSE
	RESERVE_WORDS["and"] = AND
	RESERVE_WORDS["class"] = CLASS
	RESERVE_WORDS["else"] = ELSE
	RESERVE_WORDS["false"] = FALSE
	RESERVE_WORDS["for"] = FOR
	RESERVE_WORDS["fun"] = FUN
	RESERVE_WORDS["if"] = IF
	RESERVE_WORDS["nil"] = NIL
	RESERVE_WORDS["or"] = OR
	RESERVE_WORDS["print"] = PRINT
	RESERVE_WORDS["return"] = RETURN
	RESERVE_WORDS["super"] = SUPER
	RESERVE_WORDS["this"] = THIS
	RESERVE_WORDS["true"] = TRUE
	RESERVE_WORDS["var"] = VAR
	RESERVE_WORDS["while"] = WHILE
}
func (s *Scanner) parseIdentifier() {
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
			s.AddToken(IDENTIFIER)
		}

		break
	}
}
