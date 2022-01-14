package slack_wh

//Message https://api.slack.com/reference/messaging/payload
type Message struct {
	Text     *string `json:"text,omitempty"`
	Blocks   []Block `json:"blocks,omitempty"`
	ThreadTS *string `json:"thread_ts,omitempty"`
	Mrkdwn   *bool   `json:"mrkdwn,omitempty"`
}

func NewMessage() Message {
	return Message{}
}

func (e Message) WithText(text string) Message {
	e.Text = &text
	return e
}

func (e Message) WithBlocks(blocks ...Block) Message {
	e.Blocks = blocks
	return e
}

func (e Message) WithThreadTS(threadTS string) Message {
	e.ThreadTS = &threadTS
	return e
}

func (e Message) WithMrkdwn(mrkdwn bool) Message {
	e.Mrkdwn = &mrkdwn
	return e
}
