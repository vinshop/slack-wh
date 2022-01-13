package slack_wh

type BlockElement interface {
	blockElement()
}

type InSection interface {
	inSection()
}

type InActions interface {
	inActions()
}

type InInput interface {
	inInput()
}

type InContext interface {
	inContext()
}

type BlockStyle string

const (
	BlockDefault BlockStyle = ""
	BlockPrimary BlockStyle = "primary"
	BlockDanger  BlockStyle = "danger"
)

//ElButton : https://api.slack.com/reference/block-kit/block-elements#button
type ElButton struct {
	Type     string     `json:"type"`
	Text     Text       `json:"text"`
	ActionID string     `json:"action_id"`
	URL      string     `json:"url,omitempty"`
	Value    string     `json:"value,omitempty"`
	Style    BlockStyle `json:"style,omitempty"`
	Confirm  *Confirm   `json:"confirm,omitempty"`
}

func (ElButton) blockElement() {}
func (ElButton) inSection()    {}
func (ElButton) inActions()    {}

func NewButton(ActionID string, text Text) *ElButton {
	return &ElButton{
		Type:     "button",
		Text:     text,
		ActionID: ActionID,
	}
}

func (e *ElButton) WithUrl(url string) *ElButton {
	e.URL = url
	return e
}

func (e *ElButton) WithValue(value string) *ElButton {
	e.Value = value
	return e
}

func (e *ElButton) WithStyle(style BlockStyle) *ElButton {
	e.Style = style
	return e
}

func (e *ElButton) WithConfirm(confirm *Confirm) *ElButton {
	e.Confirm = confirm
	return e
}

//ElCheckboxGroup : https://api.slack.com/reference/block-kit/block-elements#checkboxes
type ElCheckboxGroup struct {
	Type           string   `json:"type"`
	ActionID       string   `json:"action_id"`
	Options        []Option `json:"options"`
	InitialOptions []Option `json:"initial_options,omitempty"`
	Confirm        *Confirm `json:"confirm,omitempty"`
	FocusOnLoad    bool     `json:"focus_on_load,omitempty"`
}

func (ElCheckboxGroup) blockElement() {}
func (ElCheckboxGroup) inSection()    {}
func (ElCheckboxGroup) inActions()    {}
func (ElCheckboxGroup) inInput()      {}

func NewCheckBoxGroup(actionID string, options ...Option) *ElCheckboxGroup {
	return &ElCheckboxGroup{
		Type:     "checkboxes",
		ActionID: actionID,
		Options:  options,
	}
}

func (e *ElCheckboxGroup) WithInitialOptions(v []Option) *ElCheckboxGroup {
	e.InitialOptions = v
	return e
}
func (e *ElCheckboxGroup) WithConfirm(v *Confirm) *ElCheckboxGroup {
	e.Confirm = v
	return e
}
func (e *ElCheckboxGroup) WithFocusOnLoad(v bool) *ElCheckboxGroup {
	e.FocusOnLoad = v
	return e
}

//ElDatepicker : https://api.slack.com/reference/block-kit/block-elements#datepicker
type ElDatepicker struct {
	Type        string   `json:"type"`
	ActionID    string   `json:"action_id"`
	Placeholder *Text    `json:"placeholder,omitempty"`
	InitialDate string   `json:"initial_date,omitempty"`
	Confirm     *Confirm `json:"confirm,omitempty"`
	FocusOnLoad bool     `json:"focus_on_load,omitempty"`
}

func (ElDatepicker) blockElement() {}
func (ElDatepicker) inInput()      {}
func (ElDatepicker) inActions()    {}
func (ElDatepicker) inSection()    {}

//ElImage : https://api.slack.com/reference/block-kit/block-elements#image
type ElImage struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

func (ElImage) blockElement() {}
func (ElImage) inSection()    {}
func (ElImage) inContext()    {}

//Multiselect : https://api.slack.com/reference/block-kit/block-elements#image
//#TODO

//ElOverflow : https://api.slack.com/reference/block-kit/block-elements#overflow
type ElOverflow struct {
	Type     string   `json:"type"`
	ActionID string   `json:"action_id"`
	Options  []Option `json:"options"`
	Confirm  *Confirm `json:"confirm,omitempty"`
}

func (ElOverflow) blockElement() {}
func (ElOverflow) inSection()    {}
func (ElOverflow) inActions()    {}

//ElInput : https://api.slack.com/reference/block-kit/block-elements#input
type ElInput struct {
	Type                  string          `json:"type"`
	ActionID              string          `json:"action_id"`
	Placeholder           *Text           `json:"placeholder"`
	InitialValue          string          `json:"initial_value"`
	Multiline             bool            `json:"multiline"`
	MinLength             int             `json:"min_length"`
	MaxLength             int             `json:"max_length"`
	DispatchActionConfirm *DispatchAction `json:"dispatch_action_confirm"`
	FocusOnLoad           bool            `json:"focus_on_load"`
}

func (ElInput) blockElement() {}
func (ElInput) inInput()      {}

//ElRadio : https://api.slack.com/reference/block-kit/block-elements#radio
type ElRadio struct {
	Type          string   `json:"type"`
	ActionID      string   `json:"action_id"`
	Options       []Option `json:"options"`
	InitialOption *Option  `json:"initial_option,omitempty"`
	Confirm       *Confirm `json:"confirm,omitempty"`
	FocusOnLoad   bool     `json:"focus_on_load,omitempty"`
}

func (ElRadio) blockElement() {}
func (ElRadio) inSection()    {}
func (ElRadio) inActions()    {}
func (ElRadio) inInput()      {}

//Select : https://api.slack.com/reference/block-kit/block-elements#select
//#TODO

//ElTimePicker : https://api.slack.com/reference/block-kit/block-elements#timepicker
type ElTimePicker struct {
	Type        string   `json:"type"`
	ActionID    string   `json:"action_id"`
	Placeholder *Text    `json:"placeholder,omitempty"`
	InitialItem string   `json:"initial_item,omitempty"`
	Confirm     *Confirm `json:"confirm,omitempty"`
	FocusOnLoad bool     `json:"focus_on_load,omitempty"`
}

func (ElTimePicker) blockElement() {}
func (ElTimePicker) inSection()    {}
func (ElTimePicker) inActions()    {}
func (ElTimePicker) inInput()      {}
