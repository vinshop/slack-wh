package slack_wh

type TextType string

const (
	TextPlain  TextType = "plain_text"
	TextMrkDwn TextType = "mrkdwn"
)

type Action string

const (
	EnterPressed     Action = "on_enter_pressed"
	CharacterEntered Action = "on_character_entered"
)

type ConversationType string

const (
	DirectMessage          ConversationType = "im"
	MultipartDirectMessage ConversationType = "mpim"
	PrivateChannelMessage  ConversationType = "private"
	PublicChannelMessage   ConversationType = "public"
)

//Text https://api.slack.com/reference/block-kit/composition-objects#text
type Text struct {
	Type TextType `json:"type"`
	Text string   `json:"text"`

	Emoji    *bool `json:"emoji,omitempty"`
	Verbatim *bool `json:"verbatim,omitempty"`
}

func NewText(
	t TextType, text string,
) Text {
	return Text{
		Type: t,
		Text: text,
	}
}

func (e Text) WithEmoji(emoji bool) Text {
	e.Emoji = &emoji
	return e
}

func (e Text) WithVerbatim(verbatim bool) Text {
	e.Verbatim = &verbatim
	return e
}

//Confirm https://api.slack.com/reference/block-kit/composition-objects#confirm
type Confirm struct {
	Title   Text `json:"title"`
	Text    Text `json:"text"`
	Confirm Text `json:"confirm"`
	Deny    Text `json:"deny"`

	Style *string `json:"style,omitempty"`
}

func NewConfirm(
	title Text, text Text, confirm Text, deny Text,
) Confirm {
	return Confirm{
		Title:   title,
		Text:    text,
		Confirm: confirm,
		Deny:    deny,
	}
}

func (e Confirm) WithStyle(style string) Confirm {
	e.Style = &style
	return e
}

//Option https://api.slack.com/reference/block-kit/composition-objects#option
type Option struct {
	Text  Text   `json:"text"`
	Value string `json:"value"`

	Description *Text   `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

func NewOption(
	text Text, value string,
) Option {
	return Option{
		Text:  text,
		Value: value,
	}
}

func (e Option) WithDescription(description Text) Option {
	e.Description = &description
	return e
}

func (e Option) WithURL(url string) Option {
	e.URL = &url
	return e
}

//OptionGroup https://api.slack.com/reference/block-kit/composition-objects#option_group
type OptionGroup struct {
	Label   Text     `json:"label"`
	Options []Option `json:"options"`
}

func NewOptionGroup(
	label Text, options ...Option,
) OptionGroup {
	return OptionGroup{
		Label:   label,
		Options: options,
	}
}

//DispatchAction https://api.slack.com/reference/block-kit/composition-objects#dispatch_action_config
type DispatchAction struct {
	TriggerActionsOn []Action `json:"trigger_actions_on,omitempty"`
}

func NewDispatchAction() DispatchAction {
	return DispatchAction{}
}

func (e DispatchAction) WithTriggerActionsOn(triggerActionsOn ...Action) DispatchAction {
	e.TriggerActionsOn = triggerActionsOn
	return e
}

//FilterConversation https://api.slack.com/reference/block-kit/composition-objects#filter_conversations
type FilterConversation struct {
	Include                       []ConversationType `json:"include,omitempty"`
	ExcludeExternalSharedChannels *bool              `json:"exclude_external_shared_channels,omitempty"`
	ExcludeBotUsers               *bool              `json:"exclude_bot_users,omitempty"`
}

func NewFilterConversation() FilterConversation {
	return FilterConversation{}
}

func (e FilterConversation) WithInclude(include ...ConversationType) FilterConversation {
	e.Include = include
	return e
}

func (e FilterConversation) WithExcludeExternalSharedChannels(excludeExternalSharedChannels bool) FilterConversation {
	e.ExcludeExternalSharedChannels = &excludeExternalSharedChannels
	return e
}

func (e FilterConversation) WithExcludeBotUsers(excludeBotUsers bool) FilterConversation {
	e.ExcludeBotUsers = &excludeBotUsers
	return e
}
