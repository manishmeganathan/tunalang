package syntaxtree

import (
	"testing"

	"github.com/manishmeganathan/tunalang/lexer"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: lexer.Token{Type: lexer.LET, Literal: "let"},
				Name: &Identifier{
					Token: lexer.Token{Type: lexer.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: lexer.Token{Type: lexer.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
