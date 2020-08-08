package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FEIDHM = "FEIDHM"
	BOSCA    = "BOSCA"
	CEART     = "CEART"
	FALSA    = "FALSA"
	DO       = "DO"
	EILE     = "EILE"
	TABHAIR   = "TABHAIR"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fm":     FEIDHM,
	"bosca":    BOSCA,
	"ceart":   CEART,
	"falsa":  FALSA,
	"do":     DO,
	"eile":   EILE,
	"tabhair": TABHAIR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
