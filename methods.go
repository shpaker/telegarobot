package telegarobot

import (
	// "github.com/kr/pretty"

	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

// A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	type GetMeResponse struct {
		response
		Result *User
	}
	me := &GetMeResponse{}
	data, err := bot.execApi("getMe", "")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &me)
	if err != nil {
		return nil, err
	}
	if !me.Ok {
		return nil, errors.New(me.Description)
	}
	return me.Result, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
func (bot *Bot) SendMessage(chatId, text string, options *SendMessageOptions) (*Message, error) {
	type SendMessageResponse struct {
		response
		Result *Message
	}
	sendMessageResponse := &SendMessageResponse{}
	values := url.Values{}
	values.Set("chat_id", chatId)
	values.Set("text", text)
	if options != nil {
		switch {
		case options.ReplyKeyboardMarkup != nil:
			b, _ := json.Marshal(options.ReplyKeyboardMarkup)
			values.Set("reply_markup", string(b))
		case options.ReplyKeyboardHide != nil:
			b, _ := json.Marshal(options.ReplyKeyboardHide)
			values.Set("reply_markup", string(b))
		case options.ForceReply != nil:
			b, _ := json.Marshal(options.ForceReply)
			values.Set("reply_markup", string(b))
		case options.ParseMode != "":
			values.Set("parse_mode", options.ParseMode)
		case options.DisableWebPagePreview == true:
			values.Set("disable_web_page_preview", "true")
		case options.ReplyToMessageId != 0:
			values.Set("reply_to_message_id", strconv.Itoa(options.ReplyToMessageId))
		}
	}

	data, err := bot.execApi("sendMessage", values.Encode())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &sendMessageResponse)
	if err != nil {
		return nil, err
	}
	if !sendMessageResponse.Ok {
		return nil, errors.New(sendMessageResponse.Description)
	}

	return sendMessageResponse.Result, nil
}

// Use this method to forward messages of any kind. On success, the sent Message is returned.
func (bot *Bot) ForwardMessage(chatId, fromChatId string, messageId int) (*Message, error) {
	type ForwardMessageResponse struct {
		response
		Result *Message
	}

	forwardMessageResponse := &ForwardMessageResponse{}

	values := url.Values{}
	values.Set("chat_id", chatId)
	values.Set("from_chat_id", fromChatId)
	values.Set("message_id", strconv.Itoa(messageId))

	data, err := bot.execApi("forwardMessage", values.Encode())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &forwardMessageResponse)
	if err != nil {
		return nil, err
	}

	if !forwardMessageResponse.Ok {
		return nil, errors.New(forwardMessageResponse.Description)
	}

	return forwardMessageResponse.Result, nil
}

// Use this method to receive incoming updates using long polling. An Array of Update objects is returned.
func (bot *Bot) GetUpdates(options *GetUpdatesOptions) ([]*Update, error) {
	type GetUpdatesResponse struct {
		response
		Result []*Update
	}

	getUpdatesResponse := &GetUpdatesResponse{}

	values := url.Values{}

	if options != nil {
		switch {
		case options.Offset != 0:
			values.Set("offset", strconv.Itoa(options.Offset))

		case options.Limit != 0:
			values.Set("limit", strconv.Itoa(options.Limit))

		case options.Timeout != 0:
			values.Set("timeout", strconv.Itoa(options.Timeout))
		}
	}

	data, err := bot.execApi("getUpdates", values.Encode())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &getUpdatesResponse)
	if err != nil {
		return nil, err
	}

	if !getUpdatesResponse.Ok {
		return nil, errors.New(getUpdatesResponse.Description)
	}

	return getUpdatesResponse.Result, nil

}

// Use this method to send point on the map. On success, the sent Message is returned.
func (bot *Bot) SendLocation(chatId string, latitude, longitude float32) (*Message, error) {
	type SendLocationResponse struct {
		response
		Result *Message
	}

	sendLocationResponse := &SendLocationResponse{}

	values := url.Values{}
	values.Set("chat_id", chatId)
	values.Set("latitude", fmt.Sprintf("%f", latitude))
	values.Set("longitude", fmt.Sprintf("%f", longitude))

	data, err := bot.execApi("SendLocation", values.Encode())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &sendLocationResponse)
	if err != nil {
		return nil, err
	}

	if !sendLocationResponse.Ok {
		return nil, errors.New(sendLocationResponse.Description)
	}

	return sendLocationResponse.Result, nil
}

func (bot *Bot) SendChatAction(chatId, action string) (bool, error) {
	type SendChatActionResponse struct {
		response
		Result bool
	}

	sendChatActionResponse := &SendChatActionResponse{}

	values := url.Values{}
	values.Set("chat_id", chatId)
	values.Set("action", action)

	data, err := bot.execApi("sendChatAction", values.Encode())
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &sendChatActionResponse)
	if err != nil {
		return false, err
	}

	if !sendChatActionResponse.Ok {
		return false, errors.New(sendChatActionResponse.Description)
	}

	return sendChatActionResponse.Result, nil
}

// Use this method to send photos. On success, the sent Message is returned.
func (bot *Bot) SendPhotoById(chatId, photoId string, options *SendPhotoOptions) (*Message, error) {
	type SendPhotoResponse struct {
		response
		Result *Message
	}

	sendPhotoResponse := &SendPhotoResponse{}

	values := url.Values{}
	values.Set("chat_id", chatId)
	values.Set("photo", photoId)

	if options != nil {
		switch {
		case options.Caption != "":
			values.Set("caption", options.Caption)
		case options.ReplyToMessageId != 0:
			values.Set("reply_to_message_id", strconv.Itoa(options.ReplyToMessageId))
		}
	}

	data, err := bot.execApi("sendPhoto", values.Encode())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &sendPhotoResponse)
	if err != nil {
		return nil, err
	}

	if !sendPhotoResponse.Ok {
		return nil, errors.New(sendPhotoResponse.Description)
	}

	return sendPhotoResponse.Result, nil
}

// func (bot *Bot) SendPhoto(chatId, path string, options *SendPhotoOptions) (*Message, error) {
// 	type SendPhotoResponse struct {
// 		response
// 		Result *Message
// 	}

// 	sendPhotoResponse := &SendPhotoResponse{}

// 	values := url.Values{}
// 	values.Set("chat_id", chatId)

// 	/**/
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormFile("photo", filepath.Base(path))
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = io.Copy(part, file)

// 	err = writer.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	/**/

// 	if options != nil {
// 		switch {
// 		case options.Caption != "":
// 			values.Set("caption", options.Caption)
// 		case options.ReplyToMessageId != 0:
// 			values.Set("reply_to_message_id", strconv.Itoa(options.ReplyToMessageId))
// 		}
// 	}

// 	data, err := bot.execApi("sendPhoto", values.Encode())
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(data, &sendPhotoResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !sendPhotoResponse.Ok {
// 		return nil, errors.New(sendPhotoResponse.Description)
// 	}

// 	return sendPhotoResponse.Result, nil
// }
