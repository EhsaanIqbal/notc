package ast

import (
	"testing"

	"github.com/ehsaaniqbal/notc/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "hello"},
					Value: "hello",
				},
			},
		},
	}

	if program.String() != "let x = hello;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
