package slack_wh

type Style string

const (
	Default Style = ""
	Primary Style = "primary"
	Danger  Style = "danger"
)

//Button https://api.slack.com/reference/block-kit/block-elements#button
type Button struct {
	Type     string `json:"type"`
	Text     Text   `json:"text"`
	ActionID string `json:"action_id"`

	URL     *string  `json:"url,omitempty"`
	Value   *string  `json:"value,omitempty"`
	Style   *Style   `json:"style,omitempty"`
	Confirm *Confirm `json:"confirm,omitempty"`
}

func NewButton(
	text Text, actionID string,
) Button {
	return Button{
		Type:     "button",
		Text:     text,
		ActionID: actionID,
	}
}

func (e Button) WithURL(url string) Button {
	e.URL = &url
	return e
}

func (e Button) WithValue(value string) Button {
	e.Value = &value
	return e
}

func (e Button) WithStyle(style Style) Button {
	e.Style = &style
	return e
}

func (e Button) WithConfirm(confirm Confirm) Button {
	e.Confirm = &confirm
	return e
}

func (Button) inSection() {}

func (Button) inActions() {}

func (Button) element() {}

//Checkboxes https://api.slack.com/reference/block-kit/block-elements#checkboxes
type Checkboxes struct {
	Type     string `json:"type"`
	ActionID string `json:"action_id"`
	Options  Option `json:"options"`

	InitialOptions []Option `json:"initial_options,omitempty"`
	Confirm        *Confirm `json:"confirm,omitempty"`
	FocusOnLoad    *bool    `json:"focus_on_load,omitempty"`
}

func NewCheckboxes(
	actionID string, options Option,
) Checkboxes {
	return Checkboxes{
		Type:     "checkboxes",
		ActionID: actionID,
		Options:  options,
	}
}

func (e Checkboxes) WithInitialOptions(initialOptions ...Option) Checkboxes {
	e.InitialOptions = initialOptions
	return e
}

func (e Checkboxes) WithConfirm(confirm Confirm) Checkboxes {
	e.Confirm = &confirm
	return e
}

func (e Checkboxes) WithFocusOnLoad(focusOnLoad bool) Checkboxes {
	e.FocusOnLoad = &focusOnLoad
	return e
}

func (Checkboxes) inSection() {}

func (Checkboxes) inActions() {}

func (Checkboxes) inInput() {}

func (Checkboxes) element() {}

//DatePicker https://api.slack.com/reference/block-kit/block-elements#datepicker
type DatePicker struct {
	Type     string `json:"type"`
	ActionID string `json:"action_id"`

	Placeholder *Text    `json:"placeholder,omitempty"`
	InitialDate *string  `json:"initial_date,omitempty"`
	Confirm     *Confirm `json:"confirm,omitempty"`
	FocusOnLoad *bool    `json:"focus_on_load,omitempty"`
}

func NewDatePicker(
	actionID string,
) DatePicker {
	return DatePicker{
		Type:     "datepicker",
		ActionID: actionID,
	}
}

func (e DatePicker) WithPlaceholder(placeholder Text) DatePicker {
	e.Placeholder = &placeholder
	return e
}

func (e DatePicker) WithInitialDate(initialDate string) DatePicker {
	e.InitialDate = &initialDate
	return e
}

func (e DatePicker) WithConfirm(confirm Confirm) DatePicker {
	e.Confirm = &confirm
	return e
}

func (e DatePicker) WithFocusOnLoad(focusOnLoad bool) DatePicker {
	e.FocusOnLoad = &focusOnLoad
	return e
}

func (DatePicker) inSection() {}

func (DatePicker) inActions() {}

func (DatePicker) inInput() {}

func (DatePicker) element() {}

//Image https://api.slack.com/reference/block-kit/block-elements#image
type Image struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

func NewImage(
	imageURL string, altText string,
) Image {
	return Image{
		Type:     "image",
		ImageURL: imageURL,
		AltText:  altText,
	}
}

func (Image) element() {}

func (Image) inSection() {}

func (Image) inContext() {}

//Overflow https://api.slack.com/reference/block-kit/block-elements#overflow
type Overflow struct {
	Type     string   `json:"type"`
	ActionID string   `json:"action_id"`
	Options  []Option `json:"options"`

	Confirm *Confirm `json:"confirm,omitempty"`
}

func NewOverflow(
	actionID string, options ...Option,
) Overflow {
	return Overflow{
		Type:     "overflow",
		ActionID: actionID,
		Options:  options,
	}
}

func (e Overflow) WithConfirm(confirm Confirm) Overflow {
	e.Confirm = &confirm
	return e
}

func (Overflow) element() {}

func (Overflow) inSection() {}

func (Overflow) inActions() {}

//Input https://api.slack.com/reference/block-kit/block-elements#input
type Input struct {
	Type     string `json:"type"`
	ActionID string `json:"action_id"`

	Placeholder          *Text           `json:"placeholder,omitempty"`
	InitialValue         *string         `json:"initial_value,omitempty"`
	Multiline            *bool           `json:"multiline,omitempty"`
	MinLength            *int            `json:"min_length,omitempty"`
	MaxLength            *int            `json:"max_length,omitempty"`
	DispatchActionConfig *DispatchAction `json:"dispatch_action_config,omitempty"`
	FocusOnLoad          *bool           `json:"focus_on_load,omitempty"`
}

func NewInput(
	actionID string,
) Input {
	return Input{
		Type:     "plain_text_input",
		ActionID: actionID,
	}
}

func (e Input) WithPlaceholder(placeholder Text) Input {
	e.Placeholder = &placeholder
	return e
}

func (e Input) WithInitialValue(initialValue string) Input {
	e.InitialValue = &initialValue
	return e
}

func (e Input) WithMultiline(multiline bool) Input {
	e.Multiline = &multiline
	return e
}

func (e Input) WithMinLength(minLength int) Input {
	e.MinLength = &minLength
	return e
}

func (e Input) WithMaxLength(maxLength int) Input {
	e.MaxLength = &maxLength
	return e
}

func (e Input) WithDispatchActionConfig(dispatchActionConfig DispatchAction) Input {
	e.DispatchActionConfig = &dispatchActionConfig
	return e
}

func (e Input) WithFocusOnLoad(focusOnLoad bool) Input {
	e.FocusOnLoad = &focusOnLoad
	return e
}

func (Input) element() {}

func (Input) inInput() {}

//Radio https://api.slack.com/reference/block-kit/block-elements#radio
type Radio struct {
	Type     string   `json:"type"`
	ActionID string   `json:"action_id"`
	Options  []Option `json:"options"`

	InitialOption *Option  `json:"initial_option,omitempty"`
	Confirm       *Confirm `json:"confirm,omitempty"`
	FocusOnLoad   *bool    `json:"focus_on_load,omitempty"`
}

func NewRadio(
	actionID string, options ...Option,
) Radio {
	return Radio{
		Type:     "radio_buttons",
		ActionID: actionID,
		Options:  options,
	}
}

func (e Radio) WithInitialOption(initialOption Option) Radio {
	e.InitialOption = &initialOption
	return e
}

func (e Radio) WithConfirm(confirm Confirm) Radio {
	e.Confirm = &confirm
	return e
}

func (e Radio) WithFocusOnLoad(focusOnLoad bool) Radio {
	e.FocusOnLoad = &focusOnLoad
	return e
}

func (Radio) element() {}

func (Radio) inSection() {}

func (Radio) inActions() {}

func (Radio) inInput() {}

//TimePicker https://api.slack.com/reference/block-kit/block-elements#timepicker
type TimePicker struct {
	Type     string `json:"type"`
	ActionID string `json:"action_id"`

	InitialTime *string  `json:"initial_time,omitempty"`
	Confirm     *Confirm `json:"confirm,omitempty"`
	FocusOnLoad *bool    `json:"focus_on_load,omitempty"`
}

func NewTimePicker(
	actionID string,
) TimePicker {
	return TimePicker{
		Type:     "timepicker",
		ActionID: actionID,
	}
}

func (e TimePicker) WithInitialTime(initialTime string) TimePicker {
	e.InitialTime = &initialTime
	return e
}

func (e TimePicker) WithConfirm(confirm Confirm) TimePicker {
	e.Confirm = &confirm
	return e
}

func (e TimePicker) WithFocusOnLoad(focusOnLoad bool) TimePicker {
	e.FocusOnLoad = &focusOnLoad
	return e
}

func (TimePicker) element() {}

func (TimePicker) inSection() {}

func (TimePicker) inActions() {}

func (TimePicker) inInput() {}
