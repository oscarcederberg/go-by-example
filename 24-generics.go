package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
  for i, _ := range s {
    if v == s[i] {
      return i
    }
  }
  return -1
}

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

func (list *List[T]) AllElements() []T {
  var elements []T
  for e := list.head; e != nil; e = e.next {
    elements = append(elements, e.val)
  }
  return elements
}

func main() {
  var s = []string{"foo", "bar", "zoo"}

  fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
  _ = SlicesIndex[[]string, string](s, "zoo")

  list := List[int]{}
  list.Push(10)
  list.Push(13)
  list.Push(23)
  fmt.Println("list:", list.AllElements())
}
