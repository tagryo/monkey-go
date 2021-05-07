package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser は Lexer と現在読みこみのトークンと次に読みこむトークンをもつ
type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

// Parser 構造体に引数の Lexer を設定し、トークンを読みこんでそのポインタを返す
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// 2 度呼び出し、現在読みこみトークンと次のトークンを設定する
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Parser に設定された Lexer を使って、次のトークンに進む
func (p *Parser) nextToken() {
	// 次に読みこむトークンを現在読みこみトークンに設定する
	p.curToken = p.peekToken
	// Lexer で読みこんだトークンを次のトークンに設定する
	p.peekToken = p.l.NextToken()
}

// Parser のエントリーポイント。 ASTのルートノードを返す
func (p *Parser) ParseProgram() *ast.Program {
	// AST のルートノードを作成する。Statement は空で設定する
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// EOFになるまで繰り返し Parser に設定されたトークンを読みこみステートメントノードをASTに追加する
	for p.curToken.Type != token.EOF {
		// Parser のステートメントを読みこむ
		stmt := p.parseStatement()
		// ステートメントが nil でなければ、ノードのStatements に追加する
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// Parser を次のトークンに進める
		p.nextToken()
	}
	return program
}

// 現在位置が LET なら let ステートメントを読みこむ
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// let hoge = 5;
// let ステートメントとして、変数名を読みこむ
// TODO = や 5 は読み飛ばす
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// ステートメントの変数名にはIdentifier構造体にトークンとリテラルを設定する
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
