package main

import (
  "fmt"
  "time"
)

func main() {
  channel1 := make(chan string, 1)

  go func() {
    time.Sleep(2 * time.Second)
    channel1 <- "result 1"
  }()

  select {
  case result := <- channel1:
    fmt.Println(result)
  case <- time.After(1 * time.Second):
    fmt.Println("timeout 1")
  }

  channel2 := make(chan string, 1)

  go func() {
    time.Sleep(2 * time.Second)
    channel2 <- "result 2"
  }()

  select {
  case result := <- channel2:
    fmt.Println(result)
  case <- time.After(3 * time.Second):
    fmt.Println("timeout 2")
  }
}
