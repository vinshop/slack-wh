package main

import (
	_ "embed"
	"html/template"
	"os"
)

type Field struct {
	Name     string
	NameJSON string
	Type     string
	Required bool
	Imp      []string
}

type Gen struct {
	Name   string
	Fields []Field
}

//go:embed go.tmpl
var tmpl string

func main() {
	t, err := template.New("go").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	g := &Gen{
		Name: "Text",
		Fields: []Field{
			{
				Name:     "Type",
				NameJSON: "type",
				Type:     "string",
				Required: true,
			},
		},
	}

	if err := t.Execute(os.Stdout, g); err != nil {
		panic(err)
	}

}
