package telebot

// CallbackEndpoint is an interface any element capable
// of responding to a callback `\f<unique>`.
type CallbackEndpoint interface {
	CallbackUnique() string
}

// Callback object represents a query from a callback button in an
// inline keyboard.
type Callback struct {
	ID string `json:"id"`

	// For message sent to channels, Sender may be empty
	Sender *User `json:"from"`

	// Message will be set if the button that originated the query
	// was attached to a message sent by a bot.
	Message *Message `json:"message"`

	// MessageID will be set if the button was attached to a message
	// sent via the bot in inline mode.
	MessageID string `json:"inline_message_id"`

	// Data associated with the callback button. Be aware that
	// a bad client can send arbitrary data in this field.
	Data string `json:"data"`
}

// CallbackResponse builds a response to a Callback query.
//
// See also: https://core.telegram.org/bots/api#answerCallbackQuery
type CallbackResponse struct {
	// The ID of the callback to which this is a response.
	//
	// Note: Telebot sets this field automatically!
	CallbackID string `json:"callback_query_id"`

	// Text of the notification. If not specified, nothing will be
	// shown to the user.
	Text string `json:"text,omitempty"`

	// (Optional) If true, an alert will be shown by the client instead
	// of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// (Optional) URL that will be opened by the user's client.
	// If you have created a Game and accepted the conditions via
	// @BotFather, specify the URL that opens your game.
	//
	// Note: this will only work if the query comes from a game
	// callback button. Otherwise, you may use deep-linking:
	// https://telegram.me/your_bot?start=XXXX
	URL string `json:"url,omitempty"`
}

// InlineButton represents a button displayed in the message.
type InlineButton struct {
	// Unique slagish name for this kind of button,
	// try to be as specific as possible.
	//
	// It will be used as a callback endpoint.
	Unique string `json:"unique,omitempty"`

	Text        string `json:"text"`
	URL         string `json:"url,omitempty"`
	Data        string `json:"callback_data,omitempty"`
	InlineQuery string `json:"switch_inline_query,omitempty"`
	InlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`

	Action func(*Callback) `json:"-"`
}

// CallbackUnique returns InlineButto.Unique.
func (t *InlineButton) CallbackUnique() string {
	return "\f" + t.Unique
}

// CallbackUnique returns KeyboardButton.Text.
func (t *ReplyButton) CallbackUnique() string {
	return t.Text
}
