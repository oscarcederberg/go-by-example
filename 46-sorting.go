package main

import (
	"fmt"
	"slices"
)

func main() {
	strings := []string{"c", "a", "b"}
	slices.Sort(strings)
	fmt.Println("strings:", strings)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("ints:", ints)

	sorted := slices.IsSorted(ints)
	fmt.Println("sorted:", sorted)
}
