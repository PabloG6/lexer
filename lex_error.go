package main

import "fmt"

<<<<<<< HEAD
func ReportError(line int, s string) {
=======
func ReportError(line int64, s string) {
>>>>>>> a8c20af (add alphanumeric functions for reserved keywords)
	panic(fmt.Sprintf("unknown token: %s found at line %d", s, line))
}
