package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// tokens
const (
	// UwU
	tokenComment = iota
	// for if the token is a name, i.e. a variable, a function, etc.
	tokenIden
	// any number
	tokenNumber
	// "dis iws a stwing"
	tokenString
	// iws, ish
	tokenAssign
	// pwus
	tokenAdd
	// (
	tokenLparen
	// )
	tokenRparen
)

// Base class for uwupp's Token system, defined in parser.go,
// any Token will be a derivative of this.
// Use the list of tokens also defined in parser.go (token___)
// or just make up your own integer values lol
type Token struct {
	typ  int
	text string
}

// Tokeniser will take the input and turn it into a series of tokens.
func tokenise(file string) ([]Token, error) {
	var tokens []Token
	// final file for debugging/logging
	var file2 string
	var err error = nil

	// basic file before using the actual final file
	file = `
UwU twis iws a comment?
x iws 5 UwU twis iws awso a comment?
nuzzles(x)
`

	// split file into lines
	lines := strings.Split(file, "\n")
	// idx thru lines
	for line_number := 1; line_number < len(lines); line_number++ {
		// line should be one-indexed, thankfully length of array is one-indexed to match
		line := lines[line_number]

		// stores the current word, emptied every newline
		var buffer string = ""

		// is the current line part of a comment now?
		// false-ified every newline
		is_comment := false
		_ = is_comment

		// split file into chars and idx thru chars
		for char_number := 0; char_number < len(line); char_number++ {
			char := line[char_number]
			// add char to buffer
			buffer += string(char)
			// do checks now
			if strings.Contains(buffer, "UwU ") && !is_comment {
				// this is a comment, don't parse anything for the rest of the line
				is_comment = true
			}
			// do ACTUAL checks now

			if !is_comment {
				if strings.Contains(buffer, "iws ") || strings.Contains(buffer, "ish ") {
					// remove trailing space
					buffer = buffer[0 : len(buffer)-1]
					// assignment
					tokens = append(tokens, Token{tokenAssign, buffer})
					buffer = ""
				} else if strings.Contains(buffer, "(") {
					// lparen
					tokens = append(tokens, Token{tokenLparen, buffer})
					buffer = ""
				} else if strings.Contains(buffer, ")") {
					// rparen
					tokens = append(tokens, Token{tokenRparen, buffer})
					buffer = ""
				} else if regexp.MustCompile(`^[-+]?\d*\.?\d+$`).MatchString(buffer) {
					// is a number
					tokens = append(tokens, Token{tokenNumber, buffer})
					buffer = ""
				}
			}
		}

		// add what's in the buffer as a token, should tokenise every line now
		if is_comment {
			tokens = append(tokens, Token{tokenComment, buffer})
		} else {
			tokens = append(tokens, Token{tokenIden, buffer})
		}
	}

	for i := 0; i < len(tokens); i++ {
		file2 += fmt.Sprintf("{ID: %v | Data: '%v'}\n", tokens[i].typ, tokens[i].text)
	}
	os.WriteFile("log.txt", []byte(file2), 0666)

	// -- END LOGGING

	return tokens, err
}
