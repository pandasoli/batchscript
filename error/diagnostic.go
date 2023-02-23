package error


type Diagnostic struct {
  Span TextSpan
  Msg string
}

func NewDiagnostic(span TextSpan, msg string) Diagnostic {
  return Diagnostic {
    Span: span,
    Msg: msg,
  }
}
