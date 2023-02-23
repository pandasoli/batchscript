package error


type TextSpan struct {
  Start int
  Length int
  End int
}

func NewTextSpan(start int, length int) TextSpan {
  return TextSpan {
    Start: start,
    Length: length,
    End: start + length,
  }
}

func NewTextSpanFromBounds(start int, end int) TextSpan {
  return TextSpan {
    Start: start,
    Length: end - start,
    End: end,
  }
}
