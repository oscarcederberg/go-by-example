package main

import (
  "errors"
  "fmt"
)

func f(arg int) (int, error) {
  if arg == 42 {
    return -1, errors.New("can't work with 42")
  }

  return arg + 3, nil
}

var ErrorOutOfTea = fmt.Errorf("no more tea available")
var ErrorPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
  if arg == 2 {
    return ErrorOutOfTea
  } else if arg == 4 {
    return fmt.Errorf("making tea: %w", ErrorPower)
  }
  return nil
}

func main() {
  for _, i := range []int{7, 42} {
    if r, e := f(i); e != nil {
      fmt.Println("f failed:", e)
    } else {
      fmt.Println("f worked:", r)
    }
  }

  for i := range 5 {
    if e := makeTea(i); e != nil {
      if errors.Is(e, ErrorOutOfTea) {
        fmt.Println("we should buy new tea")
      } else if errors.Is(e, ErrorPower) {
        fmt.Println("now it is dark")
      } else {
        fmt.Printf("unknown error: %s\n", e)
      }
      continue
    }

    fmt.Println("tea is ready")
  }
}

