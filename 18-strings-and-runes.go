
package main

import (
  "fmt"
  "unicode/utf8"
)

func examineRune(rune rune) {
  if rune == 't' {
    fmt.Println("found tee")
  } else if rune == 'ส' {
    fmt.Println("found so sua")
  }
}

func main() {
  const string = "สวัสดี"

  fmt.Println("len:", len(string))

  for i := 0; i < len(string); i++ {
    fmt.Printf("%x ", string[i])
  }
  fmt.Println()

  fmt.Println("rune count:", utf8.RuneCountInString(string))

  for index, runeValue := range string {
    fmt.Printf("%#U starts at %d\n", runeValue, index)
  }
  fmt.Println()

  fmt.Println("using DecodeRuneInString")
  for index, width := 0, 0; index < len(string); index += width {
    runeValue, runeWidth := utf8.DecodeRuneInString(string[index:])
    fmt.Printf("%#U starts at %d\n", runeValue, index)
    width = runeWidth
    
    examineRune(runeValue)
  }
}
