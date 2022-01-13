package slack_wh

type Message struct {
	Text        string        `json:"text"`
	Blocks      []interface{} `json:"blocks,omitempty"`
	Attachments []interface{} `json:"attachments,omitempty"`
	ThreadTS    string        `json:"threadTS,omitempty"`
	Mrkdwn      bool          `json:"mrkdwn,omitempty"`
}
