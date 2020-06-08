package lexer

import "github.com/tsurusekazuki/go-interpreter/token"

type lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func New(input string) (l *lexer) {
	l = &lexer{input: input}
	l.readChar()
	return
}

func (l *lexer) readChar() {
	// 入力が終端に到着したかチェック
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// 次の文字をセット
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *lexer) NextToken() (tok token.Token) {
	switch l.ch {
	//default:
	//	if isLetter(l.ch) {
	//		tok.Literal = l.readIdentifier()
	//		return
	//	} else {
	//		tok = newToken(token.ILLEGAL, l.ch)
	//	}
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
//
//func (l *lexer) readIdentifier() string {
//	position := l.position
//	for isLetter(l.ch) {
//		l.readChar()
//	}
//	return l.input[position:l.position]
//}
//
//// 与えられた引数が英字かどうかを判別する
//func isLetter(ch byte) bool {
//	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'z' || ch <= '_'
//}