package slack_wh

type Block interface {
	block()
}

//Action https://api.slack.com/reference/block-kit/blocks#actions
type Action struct {
	Type     string         `json:"type"`
	Elements []BlockElement `json:"elements"`
	BlockID  string         `json:"block_id,omitempty"`
}

func (Action) block() {}

//Context https://api.slack.com/reference/block-kit/blocks#context
type Context struct {
	Type     string      `json:"type"`
	Elements []InContext `json:"elements"`
	BlockID  string      `json:"block_id"`
}

func (Context) block() {}

//Divider https://api.slack.com/reference/block-kit/blocks#divider
type Divider struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id"`
}

func (Divider) block() {}

//File https://api.slack.com/reference/block-kit/blocks#file
type File struct {
	Type       string `json:"type"`
	ExternalID string `json:"external_id"`
	Source     string `json:"source"`
	BlockID    string `json:"block_id,omitempty"`
}

func (File) block() {}

//Header https://api.slack.com/reference/block-kit/blocks#header
type Header struct {
	Type    string `json:"type"`
	Text    Text   `json:"text"`
	BlockID string `json:"block_id,omitempty"`
}

func (Header) block() {}

//Image https://api.slack.com/reference/block-kit/blocks#image
type Image struct {
	Type     string `json:"type"`
	ImageUrl string `json:"image_url"`
	AltText  string `json:"alt_text"`
	Title    Text   `json:"title,omitempty"`
	BlockID  string `json:"block_id,omitempty"`
}

func (Image) block() {}

//Input https://api.slack.com/reference/block-kit/blocks#input
type Input struct {
	Type           string  `json:"type"`
	Label          Text    `json:"label"`
	Element        InInput `json:"element"`
	DispatchAction bool    `json:"dispatch_action,omitempty"`
	BlockID        string  `json:"block_id,omitempty"`
	Hint           Text    `json:"hint,omitempty"`
	Optional       bool    `json:"optional"`
}

func (Input) block() {}

//Section https://api.slack.com/reference/block-kit/blocks#section
type Section struct {
	Type      string       `json:"type"`
	Text      *Text        `json:"text"`
	BlockID   string       `json:"block_id,omitempty"`
	Fields    []Text       `json:"fields,omitempty"`
	Accessory BlockElement `json:"accessory,omitempty"`
}

func (Section) block() {}
