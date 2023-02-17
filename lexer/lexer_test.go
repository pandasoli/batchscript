package lexer

import (
  "testing"
  "batchscript/token"
)

func TestNext(t *testing.T) {
  input := `set /a math=12 * 2

`

  tests := [] struct {
    kind token.TokenKind
    literal string
  } {
    {token.Identifier, "set"},
    {token.Space, " "},
    {token.Slash, "/"},
    {token.Identifier, "a"},
    {token.Space, " "},
    {token.Identifier, "math"},
    {token.Assign, "="},
    {token.Number, "12"},
    {token.Space, " "},
    {token.Star, "*"},
    {token.Space, " "},
    {token.Number, "2"},
    {token.NewLine, "\n"},
    {token.EOF, ""},
  }

  lexer := New(input)

  for i, expected := range tests {
    tok := lexer.Lex()

    if tok.Kind != expected.kind {
      t.Fatalf(
        "test %d - wrong kind.\nExpected %q, got %q.\n%+v",
        i,
        expected.kind,
        tok.Kind,
        tok,
      )
    }

    if tok.Literal != expected.literal {
      t.Fatalf(
        "test %d - wrong literal.\nExpected %q, got %q.\n%+v",
        i,
        expected.literal,
        tok.Literal,
        tok,
      )
    }
  }
}