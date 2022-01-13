package main

import (
	_ "embed"
	"html/template"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

type Field struct {
	Name      string
	NameJSON  string
	NameInput string
	Type      string
	Const     string
}

func F(name string, t string) Field {
	return Field{
		Name:      name,
		NameJSON:  snake(name),
		NameInput: lower(name),
		Type:      t,
	}
}
func (f Field) Default(s string) Field {
	f.Const = s
	return f
}

type Struct struct {
	Doc       string
	Name      string
	Required  []Field
	Optional  []Field
	Implement []string
}

type File struct {
	Package string
	Struct  []Struct
}

//go:embed go.tmpl
var tmpl string

func main() {
	t, err := template.New("go").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	composition := File{
		Package: "composite",
		Struct: []Struct{
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#text",
				Name: "Text",
				Required: []Field{
					F("Type", "string").Default("plaint_text"),
					F("Text", "string"),
				},
				Optional: []Field{
					F("Emoji", "bool"),
					F("Verbatim", "bool"),
				},
				Implement: nil,
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#confirm",
				Name: "Confirm",
				Required: []Field{
					F("Title", "Text"),
					F("Text", "Text"),
					F("Confirm", "Text"),
					F("Deny", "Text"),
				},
				Optional: []Field{
					F("Style", "string"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#option",
				Name: "Option",
				Required: []Field{
					F("Text", "Text"),
					F("Value", "string"),
				},
				Optional: []Field{
					F("Description", "*Text"),
					F("URL", "string"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#option_group",
				Name: "OptionGroup",
				Required: []Field{
					F("Label", "Text"),
					F("Options", "[]Option"),
				},
				Optional: []Field{

				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#dispatch_action_config",
				Name: "DispatchAction",
				Required: []Field{

				},
				Optional: []Field{
					F("TriggerActionsOn", "[]string"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#filter_conversations",
				Name: "FilterConversation",
				Required: []Field{

				},
				Optional: []Field{
					F("Include", "[]string"),
					F("ExcludeExternalSharedChannels", "bool"),
					F("ExcludeBotUsers", "bool"),
				},
			},
		},
	}
	os.MkdirAll("tools/composition", fs.ModeDir)
	f, err := os.Create("tools/composition/" + composition.Package + ".go")
	if err != nil {
		panic(err)
	}

	if err := t.Execute(f, composition); err != nil {
		panic(err)
	}

}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func snake(str string) string {

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func lower(s string) string {
	if s == "Type" {
		return "t"
	}
	if s == "URL" {
		return "url"
	}
	if s == "ID" {
		return "id"
	}
	return strings.ToLower(s[:1]) + s[1:]
}
