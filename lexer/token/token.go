package token


const (
  Illegal = "Illegal"
  EOF = "EOF"

  // Spaces
  NewLine = "NewLine"
  Space = "Space"

  // Literals
  Text = "Text"

  // Operators
  Assign = "Assign"
  Plus = "Plus"
  Minus = "Minus"
  Star = "Star"
  Slash = "Slash"

  // Delimiters
  LeftParen = "LeftParen"
  RightParen = "RightParen"
)

type TokenKind string
type Token struct {
  Kind     TokenKind
  Literal  string
  Position int
  Length   int
}

func New(kind TokenKind, literal string, position int, length int) Token {
  return Token {
    Kind: kind,
    Literal: literal,
    Position: position,
    Length: length,
  }
}
