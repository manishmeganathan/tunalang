package lexer

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Arithmetic Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Logical Operators
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Language keyword mapper
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// A type alias that represents the type of a token
type TokenType string

// A structure that represents a Lexilogical token
type Token struct {
	Type    TokenType
	Literal string
}

// A constructor function that generates and returns a new
// Token object for a given token type and a character byte
func NewToken(tokenType TokenType, ch byte) Token {
	// Generate and return the character as a token object
	return Token{Type: tokenType, Literal: string(ch)}
}

// A function that looks up the token type for
// a given identifier and returns it.
func LookUpIndentifier(ident string) TokenType {
	// Retrieve the token type for the
	// ident from the map of keywords
	if tok, ok := keywords[ident]; ok {
		// Return the user defined identifier
		return tok
	}

	// Return the reserved identifier
	return IDENT
}
