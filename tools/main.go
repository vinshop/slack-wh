package main

import (
	_ "embed"
	"html/template"
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

func FS(s string) Field {
	ss := strings.Split(s, ",")
	f := F(ss[0], ss[1])
	if len(ss) == 3 {
		f = f.Default(ss[2])
	}
	return f
}

func FSS(ss ...string) []Field {
	fs := make([]Field, 0, len(ss))
	for _, s := range ss {
		fs = append(fs, FS(s))
	}
	return fs
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

func genComposition(t *template.Template) {
	composition := File{
		Package: "slack_wh",
		Struct: []Struct{
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#text",
				Name: "Text",
				Required: []Field{
					F("Type", "string").Default("plain_text"),
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
					F("Description", "Text"),
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
	f, err := os.Create("composition.go")
	if err != nil {
		panic(err)
	}

	if err := t.Execute(f, composition); err != nil {
		panic(err)
	}
}

func genBlockElement(t *template.Template) {
	blockElements := File{
		Package: "slack_wh",
		Struct: []Struct{
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#button",
				Name: "Button",
				Required: []Field{
					F("Type", "string").Default("button"),
					F("Text", "Text"),
					F("ActionID", "string"),
				},
				Optional: []Field{
					F("URL", "string"),
					F("Value", "string"),
					F("Style", "string"),
					F("Confirm", "Confirm"),
				},
				Implement: []string{"inSection", "inActions", "inBlock"},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#checkboxes",
				Name: "Checkboxes",
				Required: []Field{
					F("Type", "string").Default("checkboxes"),
					F("ActionID", "string"),
					F("Options", "Option"),
				},
				Optional: []Field{
					F("InitialOptions", "[]Option"),
					F("Confirm", "Confirm"),
					F("FocusOnLoad", "bool"),
				},
				Implement: []string{"inSection", "inActions", "inInput", "inBlock"},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#datepicker",
				Name: "DatePicker",
				Implement: []string{
					"inSection", "inActions", "inInput", "inBlock",
				},
				Required: []Field{
					F("Type", "string").Default("datepicker"),
					F("ActionID", "string"),
				},
				Optional: []Field{
					F("Placeholder", "Text"),
					F("InitialDate", "string"),
					F("Confirm", "Confirm"),
					F("FocusOnLoad", "bool"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#image",
				Name: "Image",
				Implement: []string{
					"inBlock", "inSection", "inContext",
				},
				Required: []Field{
					F("Type", "string").Default("image"),
					F("ImageURL", "string"),
					F("AltText", "string"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#overflow",
				Name: "Overflow",
				Implement: []string{
					"inBlock", "inSection", "inActions",
				},
				Required: []Field{
					F("Type", "string").Default("overflow"),
					F("ActionID", "string"),
					F("Options", "[]Option"),
				},
				Optional: []Field{
					F("Confirm", "Confirm"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#input",
				Name: "Input",
				Implement: []string{
					"inBlock", "inInput",
				},
				Required: []Field{
					FS("Type,string,plain_text_input"),
					FS("ActionID,string"),
				},
				Optional: []Field{
					FS("Placeholder,Text"),
					FS("InitialValue,string"),
					FS("Multiline,bool"),
					FS("MinLength,int"),
					FS("MaxLength,int"),
					FS("DispatchActionConfig,DispatchAction"),
					FS("FocusOnLoad,bool"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#radio",
				Name: "Radio",
				Implement: []string{
					"inBlock", "inSection", "inActions", "inInput",
				},
				Required: FSS(
					"Type,string,radio_buttons",
					"ActionID,string",
					"Options,[]Option",
				),
				Optional: FSS(
					"InitialOption,Option",
					"Confirm,Confirm",
					"FocusOnLoad,bool",
				),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#timepicker",
				Name: "TimePicker",
				Implement: []string{
					"inBlock", "inSection", "inActions", "inInput",
				},
				Required: FSS(
					"Type,string,timepicker",
					"ActionID,string",
				),
				Optional: FSS(
					"InitialTime,string",
					"Confirm,Confirm",
					"FocusOnLoad,bool",
				),
			},
		},
	}
	f, err := os.Create("block_elements.go")
	if err != nil {
		panic(err)
	}

	if err := t.Execute(f, blockElements); err != nil {
		panic(err)
	}
}

func genBlock(t *template.Template) {
	blockElements := File{
		Package: "slack_wh",
		Struct: []Struct{
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#actions",
				Name: "Actions",
				Implement: []string{
					"block",
				},
				Required: FSS(
					"Type,string,actions",
					"Elements,[]InActions",
				),
				Optional: FSS(
					"BlockID,string",
				),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#context",
				Name: "Context",
				Implement: []string{
					"block",
				},
				Required: FSS(
					"Type,string,context",
					"Elements,[]InContext",
				),
				Optional: FSS(
					"BlockID,string",
				),
			},
			{
				Doc:      "https://api.slack.com/reference/block-kit/blocks#divider",
				Name:     "Divider",
				Required: FSS("Type,string,divier"),
				Optional: FSS("BlockID,string"),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#file",
				Name: "File",
				Required: FSS(
					"Type,string,file",
					"ExternalID,string",
					"Source,string",
				),
				Optional: FSS("BlockID,string"),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#header",
				Name: "Header",
				Required: FSS(
					"Type,string,header",
					"Text,Text",
				),
				Optional: FSS("BlockID,string"),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#image",
				Name: "ImageBlock",
				Required: FSS(
					"Type,string,image",
					"ImageURL,string",
					"AltText,string",
				),
				Optional: FSS(
					"Title,Text",
					"BlockID,string",
				),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#input",
				Name: "InputBlock",
				Required: FSS(
					"Type,string,input",
					"Label,Text",
					"Element,InInput",
				),
				Optional: FSS(
					"DispatchAction,bool",
					"BlockID,string",
					"Hint,Text",
					"Optional,bool",
				),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#section",
				Name: "Section",
				Required: FSS(
					"Type,string,section",
				),
				Optional: FSS(
					"Text,Text",
					"BlockID,string",
					"Fields,[]Text",
					"Accessory,Element",
				),
			},
		},
	}
	f, err := os.Create("block.go")
	if err != nil {
		panic(err)
	}

	if err := t.Execute(f, blockElements); err != nil {
		panic(err)
	}
}

func main() {
	t, err := template.New("go").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	genComposition(t)
	genBlockElement(t)
	genBlock(t)

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
