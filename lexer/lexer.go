package lexer

import (
	"batchscript/error/diagnosticbag"
	"batchscript/lexer/token"
	"strings"

	"bytes"
)


type Lexer struct {
  text string
  position int
  nextPosition int
  current byte
  Diagnostics diagnosticbag.DiagnosticBag
}

func New(text string) *Lexer {
  lexer := &Lexer {
    text: text,
    Diagnostics: diagnosticbag.New(nil),
  }

  return lexer
}


func isText(ch byte) bool {
  return !strings.Contains("+-=*/() \n", string(ch))
}


func (self *Lexer) next() *byte {
  if self.nextPosition >= len(self.text) {
    self.current = 0
  } else {
    self.current = self.text[self.nextPosition]
  }

  self.position = self.nextPosition
  self.nextPosition += 1

  return &self.current
}

func (self *Lexer) peek() byte {
  if self.nextPosition >= len(self.text) {
    return 0
  }

  return self.text[self.nextPosition]
}

func (self *Lexer) read(buffer *bytes.Buffer, test func(byte)bool) {
  for test(self.peek()) {
    buffer.WriteByte(*self.next())
  }
}


func (self *Lexer) Lex() token.Token {
  self.next()

  buffer := bytes.NewBufferString(string(self.current))
  position := self.position
  var kind token.TokenKind = token.Illegal

  switch self.current {
    case 0: kind = token.EOF

    case ' ':
      self.read(buffer, func (ch byte) bool { return ch == ' ' })
      kind = token.Space
      break
    case '\n': kind = token.NewLine

    case '+': kind = token.Plus
    case '-': kind = token.Minus
    case '=': kind = token.Assign
    case '*': kind = token.Star
    case '/': kind = token.Slash

    case '(': kind = token.LeftParen
    case ')': kind = token.RightParen

    default:
      self.read(buffer, isText)
      kind = token.Text
  }

  return token.New(kind, buffer.String(), position, buffer.Len() - 1)
}
