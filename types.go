package telegarobot

// This object represents a message.
type Message struct {
	MessageId      int      `json:"message_id"`       // Unique message identifier
	From           *User    `json:"from"`             // Optional. Sender, can be empty for messages sent to channels
	Date           int      `json:"date"`             // Date the message was sent in Unix time
	Chat           *Chat    `json:"chat"`             // Conversation the message belongs to
	ForwardFrom    *User    `json:"forward_from"`     // Optional. For forwarded messages, sender of the original message
	ForwardDate    int      `json:"forward_date"`     // Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage *Message `json:"reply_to_message"` // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	Text           string   `json:"text"`             // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	// audio 	Audio 	Optional. Message is an audio file, information about the file
	// document 	Document 	Optional. Message is a general file, information about the file
	// photo 	Array of PhotoSize 	Optional. Message is a photo, available sizes of the photo
	// sticker 	Sticker 	Optional. Message is a sticker, information about the sticker
	// video 	Video 	Optional. Message is a video, information about the video
	// voice 	Voice 	Optional. Message is a voice message, information about the file
	Caption string `json:"caption"` // Optional. Caption for the photo or video, 0-200 characters
	// contact 	Contact 	Optional. Message is a shared contact, information about the contact
	// location 	Location 	Optional. Message is a shared location, information about the location
	NewChatParticipant  *User  `json:"new_chat_participant"`  // Optional. A new member was added to the group, information about them (this member may be the bot itself)
	LeftChatParticipant *User  `json:"left_chat_participant"` // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle        string `json:"new_chat_title"`        // Optional. A chat title was changed to this value
	// new_chat_photo 	Array of PhotoSize 	Optional. A chat photo was change to this value
	DeleteChatPhoto       bool `json:"delete_chat_photo"`       // Optional. Service message: the chat photo was deleted
	GroupChatCreated      bool `json:"group_chat_created"`      // Optional. Service message: the group has been created
	SupergroupChatCreated bool `json:"supergroup_chat_created"` // Optional. Service message: the supergroup has been created
	ChannelChatCreated    bool `json:"channel_chat_created"`    // Optional. Service message: the channel has been created
	MigrateToChatId       int  `json:"migrate_to_chat_id"`      // Optional. The group has been migrated to a supergroup with the specified identifier, not exceeding 1e13 by absolute value
	MigrateFromChatId     int  `json:"migrate_from_chat_id"`    // Optional. The supergroup has been migrated from a group with the specified identifier, not exceeding 1e13 by absolute value
}

// This object represents a Telegram user or bot.
type User struct {
	Id        int    `json:"id"`         // Unique identifier for this user or bot
	FirstName string `json:"first_name"` // User‘s or bot’s first name
	LastName  string `json:"last_name"`  // Optional. User‘s or bot’s last name
	Username  string `json:"username"`   // Optional. User‘s or bot’s username
}

// This object represents an incoming update.
// Only one of the optional parameters can be present in any given update.
type Update struct {
	UpdateId int      `json:"update_id"` // The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order.
	Message  *Message `json:"message"`   // Optional. New incoming message of any kind — text, photo, sticker, etc.
	// InlineQuery        *InlineQuery        `json:"inline_query"`         // Optional. New incoming inline query
	// ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
}

// This object represents a chat.
type Chat struct {
	Id        int    `json:"update_id"`  // Unique identifier for this chat, not exceeding 1e13 by absolute value
	Type      string `json:"type"`       // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title     string `json:"title"`      // Optional. Title, for channels and group chats
	Username  string `json:"username"`   // Optional. Username, for private chats and channels if available
	FirstName string `json:"first_name"` // Optional. First name of the other party in a private chat
	LastName  string `json:"last_name"`  // Optional. Last name of the other party in a private chat
}
