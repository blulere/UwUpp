package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

// tokens
const (
	COMMENT = iota
	IDEN
	STRING
	ASSIGN
	ADD
	LPAREN
	RPAREN
)

type item struct {
	typ  int
	text string
}

// The parser is actually a lexer.
// Tokenisation is handled by Go's built in Scanner library, thank you Go
func tokenise(file string) ([]string, error) {
	var tokens []string
	var err error = nil

	// overwrite file for simplicity
	file = "x iws 5"

	var file2 = strings.NewReader(file)
	var stream scanner.Scanner
	stream.Init(file2)

	var buffer string

	// for initialise token as first token in stream,
	// as long as not end of file,
	// scan next token to iterate
	for t := stream.Scan(); t != scanner.EOF; t = stream.Scan() {

	}

	// log
	os.WriteFile("log.txt", []byte(buffer), 0666)

	fmt.Println("this is never getting finished man")

	return tokens, err
}
