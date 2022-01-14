package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"html/template"
	"os"
	"regexp"
	"strings"
)

//go:embed go.tmpl
var tmpl string

func generate(t *template.Template, name string, data interface{}) {
	f, err := os.Create(name + ".go")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	if _, err := f.Write(p); err != nil {
		panic(err)
	}
}

func genComposition(t *template.Template) {
	compositions := File{
		Package: "slack_wh",
		Type: TT(
			T("TextType,string", "TextPlain,plain_text", "TextMrkDwn,mrkdwn"),
			T("Action,string", "EnterPressed,on_enter_pressed", "CharacterEntered,on_character_entered"),
			T("ConversationType,string", "DirectMessage,im", "MultipartDirectMessage,mpim", "PrivateChannelMessage,private", "PublicChannelMessage,public"),
		),
		Struct: []Struct{
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#text",
				Name: "Text",
				Required: []Field{
					F("Type", "TextType"),
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
					F("TriggerActionsOn", "[]Action"),
				},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/composition-objects#filter_conversations",
				Name: "FilterConversation",
				Required: []Field{

				},
				Optional: []Field{
					F("Include", "[]ConversationType"),
					F("ExcludeExternalSharedChannels", "bool"),
					F("ExcludeBotUsers", "bool"),
				},
			},
		},
	}
	generate(t, "composition", compositions)
}

func genBlockElement(t *template.Template) {
	blockElements := File{
		Package: "slack_wh",
		Type: TT(
			T("Style,string", "Default,", "Primary,primary", "Danger,danger"),
		),
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
					F("Style", "Style"),
					F("Confirm", "Confirm"),
				},
				Implement: []string{"inSection", "inActions", "element"},
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
				Implement: []string{"inSection", "inActions", "inInput", "element"},
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/block-elements#datepicker",
				Name: "DatePicker",
				Implement: []string{
					"inSection", "inActions", "inInput", "element",
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
					"element", "inSection", "inContext",
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
					"element", "inSection", "inActions",
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
					"element", "inInput",
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
					"element", "inSection", "inActions", "inInput",
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
					"element", "inSection", "inActions", "inInput",
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
	generate(t, "element", blockElements)
}

func genBlock(t *template.Template) {
	blocks := File{
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
				Doc:  "https://api.slack.com/reference/block-kit/blocks#divider",
				Name: "Divider",
				Implement: []string{
					"block",
				},
				Required: FSS("Type,string,divider"),
				Optional: FSS("BlockID,string"),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#file",
				Name: "File",
				Implement: []string{
					"block",
				},
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
				Implement: []string{
					"block",
				},
				Required: FSS(
					"Type,string,header",
					"Text,Text",
				),
				Optional: FSS("BlockID,string"),
			},
			{
				Doc:  "https://api.slack.com/reference/block-kit/blocks#image",
				Name: "ImageBlock",
				Implement: []string{
					"block",
				},
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
				Implement: []string{
					"block",
				},
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
				Implement: []string{
					"block",
				},
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
	generate(t, "block", blocks)
}

func genMessage(t *template.Template) {
	message := File{
		Package: "slack_wh",
		Struct: []Struct{
			{
				Doc:  "https://api.slack.com/reference/messaging/payload",
				Name: "Message",
				Optional: FSS(
					"Text,string",
					"Blocks,[]Block",
					"ThreadTS,string",
					"Mrkdwn,bool",
				),
			},
		},
	}
	generate(t, "message", message)
}

func main() {
	t, err := template.New("go").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	genComposition(t)
	genBlockElement(t)
	genBlock(t)
	genMessage(t)

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
