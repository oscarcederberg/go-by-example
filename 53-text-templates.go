package main

import (
	"os"
	"text/template"
)

func main() {
	template1 := template.New("template1")
	template1, error := template1.Parse("value is {{.}}\n")
	if error != nil {
		panic(error)
	}

	template1 = template.Must(template1.Parse("value: {{.}}\n"))

	template1.Execute(os.Stdout, "some text")
	template1.Execute(os.Stdout, 5)
	template1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name string, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	template2 := Create("template2", "Name: {{.Name}}\n")

	template2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	template2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	template3 := Create(
		"template3", "{{if . -}} yes {{else -}} no {{end}}\n",
	)
	template3.Execute(os.Stdout, "not empty")
	template3.Execute(os.Stdout, "")

	template4 := Create(
		"template4", "range: {{range .}}{{.}} {{end}}\n",
	)
	template4.Execute(
		os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
