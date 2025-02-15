package token

type TokenType string

type Token struct{
    Type string
    Literal string
}


const (
	ILLEGAL = "ILLEGAL"  // TOKEN THAT WE DONT KNOW
	EOF = "EOF" // END OF FILE

	IDENT = "IDENT" 
	INT = "INT"

	//operators
	ASSIGN = "=" 
	PLUS = "+"

	COMMA = "," 
    SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)


var keywords = map[string]TokenType{
	"func": FUNCTION ,
	"let": LET , 
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}