package main

import "fmt"

func main() {
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range(nums) {
    sum += num
  }
  fmt.Println("sum:", sum)

  for index, _ := range(nums) {
    fmt.Println("index:", index)
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}

  for key, value := range(kvs) {
    fmt.Println(key, "->", value)
  }

  for key := range(kvs) {
    fmt.Println("key:", key)
  }

  for index, char := range("go") {
    fmt.Println(index, char)
  }
}
