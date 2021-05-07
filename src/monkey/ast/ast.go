package ast

import "monkey/token"

// Node インターフェースは TokenLiteral メソッドを実装している
type Node interface {
	TokenLiteral() string
}

// Statement インターフェースは Node インターフェースの値であり、 statementNode を実装している
type Statement interface {
	Node
	statementNode()
}

// Expression インターフェースは Node インターフェースの値であり、 expressionNode を実装している
type Expression interface {
	Node
	expressionNode()
}

// ASTのルートノード
type Program struct {
	// 複数のステートメントを持つ
	Statements []Statement
}

// Program は TokenLiteral メソッドをもつ
// Program は Node インターフェースをみたす
// TokenLiteral は、ステートメントがあれば先頭のステートメントのTokenLiteralメソッドを呼び出す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let キーワードをもつステートメントの構造体
type LetStatement struct {
	// トークンと変数名と値
	Token token.Token
	Name  *Identifier
	Value string
}

// let ステートメントは Node であり Statement でもある
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// 変数名の構造体
type Identifier struct {
	// トークンと変数名を表す値
	Token token.Token
	Value string
}

// 変数名は Node であり Expression でもある
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
