package main

// Interprets UwU++ files by calling the parser (tokeniser and lexer), then running through the syntax. Returns an error.
func interpret(file string) error {
	var tokens []Token
	var err error = nil

	// Call the tokeniser
	tokens, err = tokenise(file)
	_ = tokens
	// FIXME: what to do if error is thrown here ???

	return err
}
