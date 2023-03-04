package diagnosticbag

import (
  "batchscript/lexer/token"
  "batchscript/error/diagnostic"
  "batchscript/error/textspan"

  "fmt"
)


type DiagnosticBag struct {
  Diags []diagnostic.Diagnostic
}

func New(from *DiagnosticBag) DiagnosticBag {
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

func (self *DiagnosticBag) report(span textspan.TextSpan, msg string) {
  diag := diagnostic.New(span, msg)
  self.Diags = append(self.Diags, diag)
}

func (self *DiagnosticBag) ReportIllegalCharacter(pos int, char byte) {
  span := textspan.New(pos, 1)
  msg := "Illegal character input: '" + string(char) + "'"

  self.report(span, msg)
}

func (self *DiagnosticBag) ReportUnexpectedToken(span textspan.TextSpan, kind token.TokenKind, expected token.TokenKind) {
  msg := string("Unexpected token " + kind + ", expected " + expected)

  self.report(span, msg)
}
