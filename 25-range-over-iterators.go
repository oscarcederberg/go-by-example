package main

import (
  "fmt"
  "iter"
  "slices"
)

type List[T any] struct {
  head, tail *element[T]
}

type element[T any] struct {
  next *element[T]
  val T
}

func (list *List[T]) Push(v T) {
  if list.tail == nil {
    list.head = &element[T]{val: v}
    list.tail = list.head
  } else {
    list.tail.next = &element[T]{val: v}
    list.tail = list.tail.next
  }
}

func (list *List[T]) All() iter.Seq[T] {
  return func(yield func(T) bool) {
    for e := list.head; e != nil; e = e.next {
      if !yield(e.val) {
        return
      }
    }
  }
}

func genFib() iter.Seq[int] {
  return func(yield func(int) bool) {
    a, b := 1, 1

    for {
      if !yield(a) {
        return
      }

      a, b = b, a + b
    }
  }
}

func main() {
  list := List[int]{}
  list.Push(10)
  list.Push(13)
  list.Push(23)

  for e := range list.All() {
    fmt.Println(e)
  }

  all := slices.Collect(list.All())
  fmt.Println("all:", all)

  for n := genFib() {
    if n >= 10 {
      break
    }
    fmt.Println(n)
  }
}

