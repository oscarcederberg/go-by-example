package main

import (
  "fmt"
  "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
  for job := range jobs {
    fmt.Println("worker", id, "started job", job)
    time.Sleep(time.Second)
    fmt.Println("worker", id, "finished job", job)
    results <- job * 2
  }
}

func main() {
  const numJobs = 5
  jobs := make(chan int, numJobs)
  results := make(chan int, numJobs)

  for id := 1; id <= 3; id++ {
    go worker(id, jobs, results)
  }

  for job := 1; job <= numJobs; job++ {
    jobs <- job
  }
  close(jobs)

  for job := 1; job <= numJobs; job++ {
    <- results
  }
}
