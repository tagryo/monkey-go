package token

type TokenType string

type Token struct {
	Type	TokenType
	Literal	string
}

const (
	ILLEGAL	= "ILLEGAL"
	EOF		= "EOF"

	// 識別子 + リテラル
	IDENT	= "IDENT"	// add, foobar, x, y, ...
	INT		= "INT"		// 1343456

	// 演算子
	ASSIGN	= "="
	PLUS	= "+"
	MINUS	= "-"
	BANG	= "!"
	ASTERISK= "*"
	SLASH	= "/"

	LT	= "<"
	GT	= ">"

	EQ	= "=="
	NOT_EQ	= "!="

	// デリミタ
	COMMA	= ","
	SEMICOLON	= ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET	= "LET"
	IF	= "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType {
	"fn": 	FUNCTION,
	"let":	LET,
	"if":	IF,
	"else": ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

// 引数が正しいキーワードの場合、対応する TokenType を返す。そうでない場合 IDENT を返す
func LookupIdent(ident string) TokenType {
	if typ, ltrl := keywords[ident]; ltrl {
		return typ
	}
	return IDENT
}

