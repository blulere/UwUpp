package main

import (
	"fmt"
	"strings"
)

// to make interop between go and python easier
type Buffer struct {
	s string
}

func (b Buffer) contains(s string) bool {
	if strings.Contains(b.s, s) {
		return true
	} else {
		return false
	}
}

func (b Buffer) replace(rep string, with string) Buffer {
	b.s = strings.Replace(b.s, rep, with, -1)
	return b
}

// Add what's in the buffer to a token,
// then delete everything in the buffer.
func (b Buffer) flush(t []string) {
	t = append(t, b.s)
	b.s = ""
}

func lexer(fs string) ([]string, error) {
	var tokens []string
	var err error = nil

	// split fs by lines
	lines := strings.Split(fs, "\n")
	linesindex := 0

	// toggle for a string
	// if reached EOF and is string, something is wrong
	is_string := false

	// iterate through every line
	for linesindex < len(lines) {
		// shorthand
		line := lines[linesindex]

		// new line is not a comment by default
		is_comment := false
		// flush at the end if there is a match
		is_match := false

		// split line by chars
		var chars []byte
		charsindex := 0
		for i := 0; i < len(lines); i++ {
			chars[i] = line[i]
		}

		// for every line, initialise a buffer
		// this also flushes the buffer
		buffer := Buffer{}
		// iterate through every character
		// if there are no matches, add the character to the buffer
		for charsindex < len(chars) {
			// shorthand
			var char = chars[charsindex]

			// check if in a string
			if !is_string {
				if char == '"' {
					is_string = true
				}
			} else {
				if char == '"' {
					is_string = false
				}
			}
			// check if in a comment
			if !is_comment {
				if buffer.s == "UwU" {
					is_comment = true
				}
			} else {
				if buffer.s == "UwU" {
					is_comment = false
				}
			}

			// TODO: match at spac
			if !is_string || !is_comment {
				if buffer.s == " " {
					is_match = true
				}
			}
			// if there are no matches,
			// add the character to the buffer
			// but if there is a match,
			// flush the buffer
			if !is_match {
				buffer.s += string(char)
			} else {
				tokens = append(tokens, buffer.s+" ")
				buffer.s = ""
			}

			// increase chars index by 1
			charsindex++
		}

		// add a newline to the end of the line
		tokens = append(tokens, buffer.s+"\n")
		// increase lines index by 1
		linesindex++
		// print the buffer to stdout
		fmt.Println(line)
	}
	// return the final list of tokens + error code
	return tokens, err
}

func uwupp_to_python(fs string) (string, error) {
	var output string
	var err error = nil

	tokens, err := lexer(fs)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(tokens); i++ {
		// take each token and add it to the file
		output += tokens[i]
	}

	return output, err
}
