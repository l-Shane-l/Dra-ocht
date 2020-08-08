package lexer

import (
	"testing"

	"draíocht/token"
)

func TestNextToken(t *testing.T) {
	input := `bosca cuig = 5;
bosca deich = 10;

bosca cuir = fm(x, y) {
  x + y;
};

bosca toradh = cuir(cuig, deich);
!-/*5;
5 < 10 > 5;

do (5 < 10) {
	tabhair ceart;
} eile {
	tabhair falsa;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BOSCA, "bosca"},
		{token.IDENT, "cuig"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.BOSCA, "bosca"},
		{token.IDENT, "deich"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.BOSCA, "bosca"},
		{token.IDENT, "cuir"},
		{token.ASSIGN, "="},
		{token.FEIDHM, "fm"},
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
		{token.BOSCA, "bosca"},
		{token.IDENT, "toradh"},
		{token.ASSIGN, "="},
		{token.IDENT, "cuir"},
		{token.LPAREN, "("},
		{token.IDENT, "cuig"},
		{token.COMMA, ","},
		{token.IDENT, "deich"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.DO, "do"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.TABHAIR, "tabhair"},
		{token.CEART, "ceart"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EILE, "eile"},
		{token.LBRACE, "{"},
		{token.TABHAIR, "tabhair"},
		{token.FALSA, "falsa"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
