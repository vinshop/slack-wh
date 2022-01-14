package main

import (
	s "github.com/vinshop/slack-wh"
	"os"
)

func main() {

	wh := s.New(os.Getenv("WH_URL"))

	b := s.NewMessage().
		WithBlocks(
			s.NewHeader(s.NewText(s.TextPlain, "Budget Control")),
			s.NewDivider(),
			s.NewSection().WithText(s.NewText(s.TextPlain, "Statistic")).WithFields(
				s.NewText(s.TextPlain, "Number of source budget: 100"),
				s.NewText(s.TextPlain, "Number of campaign budget: 1000"),
			),
		)

	if err := wh.SendMessage(b); err != nil {
		panic(err)
	}

}
