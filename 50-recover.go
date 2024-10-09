package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if recovery := recover(); recovery != nil {
			fmt.Println("recovered. error:\n", recovery)
		}
	}()

	mayPanic()

	fmt.Println	("after mayPanic()")
}
