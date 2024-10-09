package main

import (
  "fmt"
  "time"
)

func main() {
  requests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests <- i
  }
  close(requests)

  limiter := time.Tick(200 * time.Millisecond)

  for request := range requests {
    <- limiter
    fmt.Println("request", request, time.Now())
  }

  burstyLimiter := make(chan time.Time, 3)
  for i := 1; i <= 3; i++ {
    burstyLimiter <- time.Now()
  }
  
  go func(){
    for tick := range time.Tick(200 * time.Millisecond) {
      burstyLimiter <- tick
    }
  }()

  burstyRequests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    burstyRequests <- i
  }
  close(burstyRequests)

  for request := range burstyRequests {
    <- burstyLimiter
    fmt.Println("request", request, time.Now())
  }
}
