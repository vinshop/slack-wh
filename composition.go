package slack_wh

type TextType string

const (
	TextPlain    TextType = "plain_text"
	TextMarkdown TextType = "mrkdwn"
)

//Text : https://api.slack.com/reference/block-kit/composition-objects#text
type Text struct {
	Type     TextType `json:"type"`
	Text     string   `json:"text"`
	Emoji    bool     `json:"emoji,omitempty"`
	Verbatim bool     `json:"verbatim,omitempty"`
}

// Confirm Dialog : https://api.slack.com/reference/block-kit/composition-objects#confirm
type Confirm struct {
	Title   Text   `json:"title"`
	Text    Text   `json:"text"`
	Confirm Text   `json:"confirm"`
	Deny    Text   `json:"deny"`
	Style   string `json:"style,omitempty"`
}

//Option : https://api.slack.com/reference/block-kit/composition-objects#option
type Option struct {
	Text        Text   `json:"text"`
	Value       string `json:"value"`
	Description *Text  `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

//GroupOption : https://api.slack.com/reference/block-kit/composition-objects#option_group
type GroupOption struct {
	Label   Text     `json:"label"`
	Options []Option `json:"options"`
}

// DispatchAction :https://api.slack.com/reference/block-kit/composition-objects#dispatch_action_config
type DispatchAction struct {
	TriggerActionsOn []string `json:"trigger_actions_on"`
}

type FilterAction string

const (
	FilterIM      FilterAction = "im"
	FilterMPIM    FilterAction = "mpim"
	FilterPrivate FilterAction = "private"
	FilterPublic  FilterAction = "public"
)

// Filter : https://api.slack.com/reference/block-kit/composition-objects#filter_conversations
type Filter struct {
	Include                       []FilterAction `json:"include,omitempty"`
	ExcludeExternalSharedChannels bool           `json:"exclude_external_shared_channels,omitempty"`
	ExcludeBotUsers               bool           `json:"exclude_bot_users"`
}
