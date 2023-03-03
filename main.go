package main

import (
  "batchscript/lexer"
  "batchscript/lexer/token"

  "bufio"
  "fmt"
  "os"
)


func main() {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Printf("> ")

    scanned := scanner.Scan()
    if !scanned { break }
    line := scanner.Text()
    if line == "" { break }

    lex := lexer.New(line)

    for {
      tok := lex.Lex()
      fmt.Printf("%+v\n", tok)

      if tok.Kind == token.EOF {
        break
      }
    }

    if len(lex.Diagnostics.Diags) > 0 {
      lex.Diagnostics.Print(line)
    }
  }
}
