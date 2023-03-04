package lexer

import (
  "testing"
  "batchscript/lexer/token"
)

func TestNext(t *testing.T) {
  input :=
`set /a math=12 * 2
echo abcde
`

  tests := [] struct {
    kind token.TokenKind
    literal string
  } {
    {token.Text, "set"},
    {token.Space, " "},
    {token.Slash, "/"},
    {token.Text, "a"},
    {token.Space, " "},
    {token.Text, "math"},
    {token.Assign, "="},
    {token.Text, "12"},
    {token.Space, " "},
    {token.Star, "*"},
    {token.Space, " "},
    {token.Text, "2"},
    {token.NewLine, "\n"},

    {token.Text, "echo"},
    {token.Space, " "},
    {token.Text, "abcde"},
    {token.NewLine, "\n"},
    {token.EOF, "\x00"},
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
