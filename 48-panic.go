package main

import "os"

func main() {
	panic("a problem")

	_, error := os.Create("/tmp/file")
	if error != nil {
		panic(error)
	}
}
