package textspan


type TextSpan struct {
  Start int
  Length int
  End int
}

func New(start int, length int) TextSpan {
  return TextSpan {
    Start: start,
    Length: length,
    End: start + length,
  }
}

func NewFromBounds(start int, end int) TextSpan {
  return TextSpan {
    Start: start,
    Length: end - start,
    End: end,
  }
}
