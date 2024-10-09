package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

func main() {
  var counter atomic.Uint64

  var waitGroup sync.WaitGroup

  for i := 0; i < 50; i++ {
    waitGroup.Add(1)

    go func(){
      for c := 0; c < 1000; c++ {
        counter.Add(1)
      }

      waitGroup.Done()
    }()
  }

  waitGroup.Wait()
  fmt.Println("counter:", counter.Load())
}
