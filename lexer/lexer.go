package lexer

// A structure that represents a Lexer
type Lexer struct {
	// Represents the input to the lexer
	input string

	// Represents the position in the input (current char)
	positionCurrent int

	// Represents the current reading position (after current char)
	positionNext int

	// Represents the current char
	ch byte
}

// A constructor function that generates and
// returns an initialised Lexer object
func NewLexer(input string) *Lexer {
	// Construct a lexer with the input
	l := &Lexer{input: input}
	// Read the first character of the input
	// to initialise the lexer
	l.ReadChar()

	// Return the lexer
	return l
}

func (l *Lexer) NextToken() Token {
	// Declare a token
	var tok Token

	// Eat all whitespaces until next character
	l.EatWhitespaces()

	// Check the value of the character read by the lexer
	switch l.ch {
	case '=':
		// Check if the next character is a '='
		if l.PeekChar() == '=' {
			// Move lexer to the next character
			l.ReadChar()
			// Set the token value to '=='
			tok = Token{Type: EQ, Literal: "=="}

		} else {
			// Set the token value to '='
			tok = NewToken(ASSIGN, l.ch)
		}
	case '!':
		// Check if the next character is a '='
		if l.PeekChar() == '=' {
			// Move lexer to the next character
			l.ReadChar()
			// Set the token value to '!='
			tok = Token{Type: NOT_EQ, Literal: "!="}

		} else {
			// Set the token value to '!'
			tok = NewToken(BANG, l.ch)
		}

	case '+':
		tok = NewToken(PLUS, l.ch)
	case '-':
		tok = NewToken(MINUS, l.ch)
	case '/':
		tok = NewToken(SLASH, l.ch)
	case '*':
		tok = NewToken(ASTERISK, l.ch)
	case '<':
		tok = NewToken(LT, l.ch)
	case '>':
		tok = NewToken(GT, l.ch)
	case ':':
		tok = NewToken(COLON, l.ch)
	case ';':
		tok = NewToken(SEMICOLON, l.ch)
	case '(':
		tok = NewToken(LPAREN, l.ch)
	case ')':
		tok = NewToken(RPAREN, l.ch)
	case ',':
		tok = NewToken(COMMA, l.ch)
	case '{':
		tok = NewToken(LBRACE, l.ch)
	case '}':
		tok = NewToken(RBRACE, l.ch)
	case '[':
		tok = NewToken(LBRACK, l.ch)
	case ']':
		tok = NewToken(RBRACK, l.ch)

	case '"':
		tok.Type = STRING
		tok.Literal = l.ReadString()
	case 0:
		// End of File
		tok.Literal = ""
		tok.Type = EOF

	default:
		// Check if character is a letter/digit
		if isLetter(l.ch) {
			// Identifier Detected - Read the full identifier
			tok.Literal = l.ReadIdentifier()
			// Get the mapping of the identifier literal to the type
			tok.Type = LookUpIndentifier(tok.Literal)
			// Return the identifier token
			return tok

		} else if isDigit(l.ch) {
			// Number Detected - Read the full number
			tok.Literal = l.ReadNumber()
			// Set the token type
			tok.Type = INT
			// Return the numeric token
			return tok

		} else {
			// Illegal Token
			tok = NewToken(ILLEGAL, l.ch)
		}
	}

	// Read the next character from the lexer input
	l.ReadChar()
	// Return the lexed token
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
