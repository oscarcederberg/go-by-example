package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	expression, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(expression.MatchString("peach"))
	fmt.Println(expression.FindString("peach punch"))
	fmt.Println("index:", expression.FindStringIndex("peach punch"))
	fmt.Println(expression.FindStringSubmatch("peach punch"))
	fmt.Println(expression.FindStringSubmatchIndex("peach punch"))
	fmt.Println(expression.FindAllString("peach punch pinch", -1))
	fmt.Println("all:", expression.FindAllStringSubmatchIndex(
		"peach punch pinch",
		-1,
	))
	fmt.Println(expression.FindAllString("peach punch pinch", 2))
	fmt.Println(expression.Match([]byte("peach")))

	expression = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", expression)
	fmt.Println(expression.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := expression.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

