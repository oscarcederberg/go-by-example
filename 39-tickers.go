package main

import (
  "fmt"
  "time"
)

func main() {
  ticker := time.NewTicker(500 * time.Millisecond)
  done := make(chan bool)

  go func() {
    for {
      select {
      case <- done:
        return
      case tick := <- ticker.C:
        fmt.Println("tick at", tick)
      }
    }
  }()

  time.Sleep(1600 * time.Millisecond)
  ticker.Stop()
  done <- true
  fmt.Println("ticker stopped")
}
