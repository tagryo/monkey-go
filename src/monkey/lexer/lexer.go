package lexer

import "monkey/token"

type Lexer struct {
	input			string
	position		int		// 入力における現在の位置（現在の文字を指し示す）
	readPosition	int		// これから読みこむ位置（現在の文字の次）
	ch				byte	// 現在検査中の文字
}

// inputを持つ Lexer 構造体のポインタを返す
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // l の初期化
	return l
}

// 引数のLexer構造体のreadPositionを読みこみ、chに設定する。存在しない場合chに0を設定する
// その後position, readPositionを調整する。
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// Lexer の「次の文字の位置」が文字列長より長い場合
		// 検査中の文字を0とする
		l.ch = 0
	} else {
		// そうでない場合（普通に次の文字がある場合）
		// 次の文字を読みこみ、検査中の文字とする
		l.ch = l.input[l.readPosition]
	}
	// いま読みこんだ位置を現在位置とする
	l.position = l.readPosition
	// 次に読みこむ位置を+1する
	l.readPosition += 1
}

// 解析器の文字を Token 型に変換し、次の位置に移動する
// 変換した Token を返す
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// 引数の解析器の読み込み中の文字を Token 型に変換する
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// 文字の場合、文字列を読みこんでLiteralに設定し、次に進む
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) 
			return tok
		} else if isDigit(l.ch) {
			// 数字の場合、数字を読みこんでLiteralに設定し、次に進む
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			// 各ケースと文字でもない場合、ILLEGALケースとする
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// 解析器が読みこむ位置を次に移動する
	l.readChar()

	// 読みこんだ文字のToken型を返す
	return tok
}

// Type に引数の tokenType を、 Literal に string キャストしたch を設定した Token 型を返す
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Lexer に「スペース文字なら次の文字を読みこむメソッド」を定義する
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Lexer に「文字を読み進めて文字列を返すメソッド」を定義する
func (l *Lexer) readIdentifier() string {
	// 開始インデックスを記憶し、文字は読み進めて、元の文字列の開始インデックスから現在位置までの文字列サブセットを返す
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Lexer に「数字を読み進めて数字文字列を返すメソッド」を定義する
func (l *Lexer) readNumber() string {
	// 開始インデックスを記憶し、文字(数字)は読み進めて、元の文字列の開始インデックスから現在位置までの文字列サブセットを返す
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Lexer に「文字を読み進めないで次の文字を返すメソッド」を定義する
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// 引数が a-zA-Z_ の場合、true を返す
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/// 引数が 0-9 の場合、true を返す
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

