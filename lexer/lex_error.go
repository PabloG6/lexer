package lexer

import "fmt"

func ReportError(line int64, s string) {
	panic(fmt.Sprintf("unknown token: %s found at line %d", s, line))
}
