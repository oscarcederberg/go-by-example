package main

import (
  "fmt"
  "time"
)

func main() {
  channel1 := make(chan string, 1)
  channel2 := make(chan string, 1)

  go func() {
    time.Sleep(1 * time.Second)
    channel1 <- "one"
  }()

  go func() {
    time.Sleep(2 * time.Second)
    channel2 <- "two"
  }()

  for i := 0; i < 2; i++ {
    select {
    case message1 := <- channel1:
      fmt.Println("received", message1)
    case message2 := <- channel2:
      fmt.Println("received", message2)
    }
  }
}
