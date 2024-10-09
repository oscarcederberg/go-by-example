package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	fmt.Println("creating")
	file, error := os.Create(path)
	if error != nil {
		panic(error)
	}
	return file
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	error := file.Close()

	if error != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", error)
		os.Exit(1)
	}
}

func main() {
	file := createFile("/tmp/defer.txt")
	defer closeFile(file)
	writeFile(file)
}
