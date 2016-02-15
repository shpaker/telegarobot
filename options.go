package telegarobot

type SendMessageOptions struct {
	ParseMode             string `json:"parse_mode"`               // Optional Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool   `json:"disable_web_page_preview"` // Optional Disables link previews for links in this message
	ReplyToMessageId      int    `json:"reply_to_message_id"`      // Optional If the message is a reply, ID of the original message
	keyboardOptions
}

type GetUpdatesOptions struct {
	Offset  int `json:"offset"`  // Optional Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id.
	Limit   int `json:"limit"`   // Optional Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100
	Timeout int `json:"timeout"` // Optional Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling
}

type SendPhotoOptions struct {
	Caption          string `json:"caption"`             // Optional Photo caption (may also be used when resending photos by file_id), 0-200 characters.
	ReplyToMessageId int    `json:"reply_to_message_id"` // Optional If the message is a reply, ID of the original message
	// reply_markup 	ReplyKeyboardMarkup or ReplyKeyboardHide or ForceReply 	Optional 	Additional interface options. A JSON-serialized object for a custom reply keyboard, instructions to hide keyboard or to force a reply from the user.
}

type keyboardOptions struct {
	ReplyKeyboardMarkup *ReplyKeyboardMarkup `json:"reply_markup"`
	ReplyKeyboardHide   *ReplyKeyboardHide   `json:"reply_markup"`
	ForceReply          *ForceReply          `json:"reply_markup"`
}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).'
type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"`          // Array of Array of String 	Array of button rows, each represented by an Array of Strings
	ResizeKeyboard  bool       `json:"resize_keyboard"`   // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard bool       `json:"one_time_keyboard"` // Optional. Requests clients to hide the keyboard as soon as it's been used. Defaults to false.
	Selective       bool       `json:"selective"`         // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// Upon receiving a message with this object, Telegram clients will hide the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"` // Requests clients to hide the custom keyboard
	Selective    bool `json:"selective"`     // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool `json:"force_reply"` // Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	Selective  bool `json:"selective"`   // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}
