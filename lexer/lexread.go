package lexer

// A method of Lexer that moves the lexer's cursor to the
// next character and skips all whitespaces in between.
func (l *Lexer) EatWhitespaces() {
	// Iterate until the read character is a whitespace
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

// A method of Lexer that reads a single character from
// the lexer input and moves lexer to the next character
func (l *Lexer) ReadChar() {
	// Check if the end of input has been reached
	if l.positionNext >= len(l.input) {
		// Assign character to 0
		l.ch = 0
	} else {
		// Assign character to next input character
		l.ch = l.input[l.positionNext]
	}

	// Move current position to the next position
	l.positionCurrent = l.positionNext
	// Increment the next position
	l.positionNext += 1
}

// A method of Lexer that reads an identifier token from the lexer input
func (l *Lexer) ReadIdentifier() string {
	// Retrieve the starting position of the identifier
	position := l.positionCurrent

	// Iterate over the input until characters are letters
	for isLetter(l.ch) {
		l.ReadChar()
	}

	// Extract the identifier from the input with the start and current position
	return l.input[position:l.positionCurrent]
}

// A method of Lexer that reads a numeric token from the lexer input
func (l *Lexer) ReadNumber() string {
	// Retrieve the starting position of the number
	position := l.positionCurrent

	// Iterate over the input until characters are digits
	for isDigit(l.ch) {
		l.ReadChar()
	}

	// Extract the number from the input with the start and current position
	return l.input[position:l.positionCurrent]
}
