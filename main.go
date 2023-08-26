package main

import "os"

var hasError bool = false

func main() {
	print("Hello world!")
	arguments := os.Args
	if len(arguments) > 1 {
		println("jlox: [script]")
		os.Exit(64)

	} else if len(arguments) == 0 {
		runFile(arguments[0])
	} else {
		runPrompt()
	}
}

func runFile(args string) {
	bytes, err := os.ReadFile(args)
	checkErr(err)

	run(bytes)

}

func run(source []byte) {
	for idx := range source {
		println("%c", source[idx])
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func runPrompt() {

}

func handleError(line int32, where string, message string) {
	println("error occurred on line %d, %s. Error message: %s", line, where, message)
	hasError = true
}

type Token struct {
}
