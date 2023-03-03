package error

import (
  "batchscript/lexer/token"

  "fmt"
)


type DiagnosticBag struct {
  Diags []Diagnostic
}

func NewDiagnosticBag(from *DiagnosticBag) DiagnosticBag {
  res := DiagnosticBag {}

  if from != nil {
    res.Diags = from.Diags
  }

  return res
}


func (self *DiagnosticBag) Print(code string) {
  for diagi := 0; diagi < len(self.Diags); diagi++ {
    diag := self.Diags[diagi]

    fmt.Println()
    fmt.Printf("\033[31m%s\033[0m\n", diag.Msg);

    prefix := code[0:diag.Span.Start]
    error  := code[diag.Span.Start:diag.Span.End]
    suffix := code[diag.Span.End:]

    fmt.Printf("  ╰─ %s", prefix)
    fmt.Printf("\033[31m%s\033[0m", error)
    fmt.Printf("%s\n", suffix)
  }

  fmt.Println()
}

func (self *DiagnosticBag) report(span TextSpan, msg string) {
  diag := NewDiagnostic(span, msg)
  self.Diags = append(self.Diags, diag)
}

func (self *DiagnosticBag) ReportIllegalCharacter(pos int, char byte) {
  span := NewTextSpan(pos, 1)
  msg := "Illegal character input: '" + string(char) + "'"

  self.report(span, msg)
}

func (self *DiagnosticBag) ReportUnexpectedToken(span TextSpan, kind token.TokenKind, expected token.TokenKind) {
  msg := string("Unexpected token " + kind + ", expected " + expected)

  self.report(span, msg)
}
