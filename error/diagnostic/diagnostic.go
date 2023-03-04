package diagnostic

import (
  "batchscript/error/textspan"
)


type Diagnostic struct {
  Span textspan.TextSpan
  Msg string
}

func New(span textspan.TextSpan, msg string) Diagnostic {
  return Diagnostic {
    Span: span,
    Msg: msg,
  }
}
