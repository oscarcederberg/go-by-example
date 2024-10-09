package main

import "fmt"

func factorial(i int) int {
  if i == 1 {
    return 1
  }
  return i * factorial(i - 1)
}

func main() {
  fmt.Println("5! =", factorial(5))

  var fibonacci func(n int) int

  fibonacci = func(n int) int {
    if n < 2 {
      return n
    }

    return fibonacci(n - 1) + fibonacci(n - 2)
  }

  fmt.Println("7 fibonacci =", fibonacci(7))
}
