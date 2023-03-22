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

		// split line by words
		words := strings.Split(line, " ")
		wordsindex := 0

		// for every line, initialise a buffer
		// this also flushes the buffer
		buffer := Buffer{}
		// iterate through every word
		for wordsindex < len(words) {
			// shorthand
			word := words[wordsindex]

			// split each keyword by space UNLESS is_comment or is_string
			if !is_comment || !is_string {
				if word == " " {
					// then flush the buffer
					buffer.flush(tokens)
				}
				// now, match the buffer to a string, if it matches, do the same thing
				if buffer.contains("UwU") {
					buffer.replace("UwU", "# UwU")
					buffer.flush(tokens)
					is_comment = true
				}

				if buffer.contains("iws") {
					buffer.replace("iws", "=")
					buffer.flush(tokens)
				}
			}

			// add current word to buffer + space
			buffer.s += word + " "

			// increase words index by 1
			wordsindex++
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
