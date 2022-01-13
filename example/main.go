package main

import (
	"encoding/json"
	"fmt"
	s "github.com/vinshop/slack-wh"
)

func main() {
	b := s.Section{
		Type: "section",
		Text: &s.Text{
			Type:     "plain_text",
			Text:     "A",
			Emoji:    false,
			Verbatim: false,
		},
		BlockID: "",
		Fields:  nil,
		Accessory: s.NewCheckBoxGroup("1", ),
	}
	data, _ := json.Marshal(b)

	fmt.Println(string(data))
}