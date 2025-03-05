package lexer

import (
	"testing"

	"github.com/boxy-pug/monkey/token"
)

func TestNextToken(t *testing.T) {

	tests := []struct {
		input          string
		expectedTokens []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}
	}{
		{
			input: "=+(){},;",
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			input: `let five = 5; 
			let ten = 10; 
			
			let add = fn(x, y) {
			  x + y;
			};

			let result = add(five, ten);
			`,
			expectedTokens: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		l := New(tt.input)

		for i, expectedToken := range tt.expectedTokens {
			tok := l.NextToken()

			if tok.Type != expectedToken.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expectedToken.expectedType, tok.Type)
			}

			if tok.Literal != expectedToken.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expectedToken.expectedLiteral, tok.Literal)
			}
		}

	}
}
