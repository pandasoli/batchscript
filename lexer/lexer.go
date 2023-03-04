package lexer

import (
  "batchscript/lexer/token"
  "batchscript/error/diagnosticbag"

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


func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
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
      if isLetter(self.current) {
        self.read(buffer, isLetter)
        kind = token.Identifier
      } else if isDigit(self.current) {
        self.read(buffer, isDigit)
        kind = token.Number
      } else {
        self.Diagnostics.ReportIllegalCharacter(position, self.current)
      }
  }

  return token.New(kind, buffer.String(), position, buffer.Len() - 1)
}
