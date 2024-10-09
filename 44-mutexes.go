package main

import (
  "fmt"
  "sync"
)

type Container struct {
  mutex sync.Mutex
  counters map[string]int
}

func (c *Container) increment(name string){
  c.mutex.Lock()
  defer c.mutex.Unlock()
  c.counters[name]++
}

func main() {
  container := Container{
    counters: map[string]int{"a": 0, "b": 0},
  }

  var waitGroup sync.WaitGroup

  doIncrement := func(name string, n int){
    for i := 0; i < n; i++ {
      container.increment(name)
    }
    waitGroup.Done()
  }

  waitGroup.Add(3)
  go doIncrement("a", 10000)
  go doIncrement("a", 10000)
  go doIncrement("b", 10000)

  waitGroup.Wait()
  fmt.Println(container.counters)
}
