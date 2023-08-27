package main

import "fmt"

func ReportError(line int, s string) {
	panic(fmt.Sprintf("unknown token: %s found at line %d", s, line))
}
