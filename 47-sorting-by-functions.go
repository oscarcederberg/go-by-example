package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{
		"peach", "banana", "kiwi", "mango", "orange", "pear", "apple",
	}
	lengthCompare := func(a string, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lengthCompare)
	fmt.Println(fruits)

	type Person struct {
		name string
		age int
	}

	people := []Person {
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(
		people,
		func(a Person, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
