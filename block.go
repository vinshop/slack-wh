package slack_wh

//Actions https://api.slack.com/reference/block-kit/blocks#actions
type Actions struct {
	Type     string      `json:"type"`
	Elements []InActions `json:"elements"`

	BlockID *string `json:"block_id,omitempty"`
}

func NewActions(
	elements ...InActions,
) Actions {
	return Actions{
		Type:     "actions",
		Elements: elements,
	}
}

func (e Actions) WithBlockID(blockID string) Actions {
	e.BlockID = &blockID
	return e
}

func (Actions) block() {}

//Context https://api.slack.com/reference/block-kit/blocks#context
type Context struct {
	Type     string      `json:"type"`
	Elements []InContext `json:"elements"`

	BlockID *string `json:"block_id,omitempty"`
}

func NewContext(
	elements ...InContext,
) Context {
	return Context{
		Type:     "context",
		Elements: elements,
	}
}

func (e Context) WithBlockID(blockID string) Context {
	e.BlockID = &blockID
	return e
}

func (Context) block() {}

//Divider https://api.slack.com/reference/block-kit/blocks#divider
type Divider struct {
	Type string `json:"type"`

	BlockID *string `json:"block_id,omitempty"`
}

func NewDivider() Divider {
	return Divider{
		Type: "divider",
	}
}

func (e Divider) WithBlockID(blockID string) Divider {
	e.BlockID = &blockID
	return e
}

func (Divider) block() {}

//File https://api.slack.com/reference/block-kit/blocks#file
type File struct {
	Type       string `json:"type"`
	ExternalID string `json:"external_id"`
	Source     string `json:"source"`

	BlockID *string `json:"block_id,omitempty"`
}

func NewFile(
	externalID string, source string,
) File {
	return File{
		Type:       "file",
		ExternalID: externalID,
		Source:     source,
	}
}

func (e File) WithBlockID(blockID string) File {
	e.BlockID = &blockID
	return e
}

func (File) block() {}

//Header https://api.slack.com/reference/block-kit/blocks#header
type Header struct {
	Type string `json:"type"`
	Text Text   `json:"text"`

	BlockID *string `json:"block_id,omitempty"`
}

func NewHeader(
	text Text,
) Header {
	return Header{
		Type: "header",
		Text: text,
	}
}

func (e Header) WithBlockID(blockID string) Header {
	e.BlockID = &blockID
	return e
}

func (Header) block() {}

//ImageBlock https://api.slack.com/reference/block-kit/blocks#image
type ImageBlock struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`

	Title   *Text   `json:"title,omitempty"`
	BlockID *string `json:"block_id,omitempty"`
}

func NewImageBlock(
	imageURL string, altText string,
) ImageBlock {
	return ImageBlock{
		Type:     "image",
		ImageURL: imageURL,
		AltText:  altText,
	}
}

func (e ImageBlock) WithTitle(title Text) ImageBlock {
	e.Title = &title
	return e
}

func (e ImageBlock) WithBlockID(blockID string) ImageBlock {
	e.BlockID = &blockID
	return e
}

func (ImageBlock) block() {}

//InputBlock https://api.slack.com/reference/block-kit/blocks#input
type InputBlock struct {
	Type    string  `json:"type"`
	Label   Text    `json:"label"`
	Element InInput `json:"element"`

	DispatchAction *bool   `json:"dispatch_action,omitempty"`
	BlockID        *string `json:"block_id,omitempty"`
	Hint           *Text   `json:"hint,omitempty"`
	Optional       *bool   `json:"optional,omitempty"`
}

func NewInputBlock(
	label Text, element InInput,
) InputBlock {
	return InputBlock{
		Type:    "input",
		Label:   label,
		Element: element,
	}
}

func (e InputBlock) WithDispatchAction(dispatchAction bool) InputBlock {
	e.DispatchAction = &dispatchAction
	return e
}

func (e InputBlock) WithBlockID(blockID string) InputBlock {
	e.BlockID = &blockID
	return e
}

func (e InputBlock) WithHint(hint Text) InputBlock {
	e.Hint = &hint
	return e
}

func (e InputBlock) WithOptional(optional bool) InputBlock {
	e.Optional = &optional
	return e
}

func (InputBlock) block() {}

//Section https://api.slack.com/reference/block-kit/blocks#section
type Section struct {
	Type string `json:"type"`

	Text      *Text    `json:"text,omitempty"`
	BlockID   *string  `json:"block_id,omitempty"`
	Fields    []Text   `json:"fields,omitempty"`
	Accessory *Element `json:"accessory,omitempty"`
}

func NewSection() Section {
	return Section{
		Type: "section",
	}
}

func (e Section) WithText(text Text) Section {
	e.Text = &text
	return e
}

func (e Section) WithBlockID(blockID string) Section {
	e.BlockID = &blockID
	return e
}

func (e Section) WithFields(fields ...Text) Section {
	e.Fields = fields
	return e
}

func (e Section) WithAccessory(accessory Element) Section {
	e.Accessory = &accessory
	return e
}

func (Section) block() {}
