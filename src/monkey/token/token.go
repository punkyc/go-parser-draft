
package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF ="EOF"

	//标识符 + 字面量
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT = "INT"

	//运算符
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	
	LT = "<"
	GT = ">"

	EQ = "=="
	NOT_EQ = "!="

	//分隔符
	COMMA = ","
	SEMICOLON = ";"
	COLON = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	STRING = "STRING"

	//关键字
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}
// 检查关键字判断给定的标识符是否是关键字，以此区分开关键字和用户定义的标识符
func LookupIdent(ident string) TokenType {
	if tok, ok :=keywords[ident]; ok {
		return tok
	}
	return IDENT
}