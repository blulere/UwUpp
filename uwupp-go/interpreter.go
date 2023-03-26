// this file keeps getting corrupted by git and i have no idea why

package main

import (
	"strings"
)

type Buffer struct {
	s string
}

func (b Buffer) replace(rep string, with string) Buffer {
	b.s = strings.Replace(b.s, rep, with, 1)
	return b
}

func parser(fs string) ([]string, error) {
	// no match: add a new char to the buffer
	// a match: clear the buffer and add a token
	// newlines are a match and do shit on every newline
	// get this shit working before adding comments and string support
	var tokens []string
	var err error = nil
	buffer := Buffer{}

	// split fs into chars
	// + add a newline at the end for parsing sanity
	chars := []byte(fs)
	chars = append(chars, byte('\n'))
	charsindex := 0

	is_string := false
	is_comment := false
	is_match := false
	// Instead of adding it to the buffer and clearing
	// it, try instead just replacing text in the
	// buffer and printing that
	// It's a bit nuclear but it should be adaptable
	// for string support
	for charsindex < len(chars) {
		char := chars[charsindex]

		charsindex++
	}

	return tokens, err
}

// we are no longer re-inventing the wheel baby
func uwupp_to_python(fs string) (string, error) {
	var pythoncode string
	var err error = nil

	tokens, err := parser(fs)

	for i := 0; i < len(tokens); i++ {
		pythoncode += tokens[i]
	}

	return pythoncode, err
}
