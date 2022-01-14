package main

import (
	"encoding/json"
	"fmt"
	s "github.com/vinshop/slack-wh"
)

func main() {

	b := s.NewActions(
		s.NewButton(s.NewText("Choose"), "choosetrue").WithStyle("success"),
		s.NewButton(s.NewText("Choose"), "choose").WithStyle("danger"),
	)

	data, _ := json.Marshal(b)

	fmt.Println(string(data))
}
