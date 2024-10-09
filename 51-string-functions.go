package main

import (
	"fmt"
	str "strings"
)

var print = fmt.Println

func main() {
	print("contains:", str.Contains("test", "es"))
	print("count:", str.Count("test", "t"))
	print("hasprefix:", str.HasPrefix("test", "te"))
	print("hassuffix:", str.HasSuffix("test", "st"))
	print("index:", str.Index("test", "e"))
	print("join:", str.Join([]string{"a", "b"}, "-"))
	print("repeat:", str.Repeat("a", 5))
	print("replace:", str.Replace("foo", "o", "0", -1))
	print("replace:", str.Replace("foo", "o", "0", 1))
	print("split:", str.Split("a-b-c-d-e", "-"))
	print("tolower:", str.ToLower("TEST"))
	print("toupper:", str.ToUpper("test"))
}
