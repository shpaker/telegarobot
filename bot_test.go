package telegarobot

import (
	"fmt"
	"testing"
)

const (
	TOKEN = "138664782:AAHovZk0XY5O-1yY-2vWvcBjQ-CQPqMANdQ"
	ID    = "9429534"
)

var b = &Bot{TOKEN}

func TestGetMe(t *testing.T) {
	_, err := b.GetMe()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestSendMessage(t *testing.T) {
	options := &SendMessageOptions{
		ParseMode:             "markdown",
		DisableWebPagePreview: true,
		ReplyToMessageId:      1,
	}
	// options.ReplyKeyboardMarkup = &ReplyKeyboardMarkup{
	// 	Keyboard:        [][]string{{"🐹", "🐰", "🐻"}, {"🐶", "🐱", "🐭"}, {"🐼", "🐨", "🐵"}},
	// 	OneTimeKeyboard: true,
	// }
	_, err := b.SendMessage("9429534",
		"_test_ *sendMessage* [link](google.com) `inline` ```outlone```",
		options)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestForwardMessage(t *testing.T) {
	_, err := b.ForwardMessage(ID, ID, 1)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestGetUpdates(t *testing.T) {
	_, err := b.GetUpdates(nil)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestSendLocation(t *testing.T) {
	_, err := b.SendLocation(ID, 81.384000, 54.918937)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestSendChatAction(t *testing.T) {
	_, err := b.SendChatAction(ID, ActionTyping)

	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestSendPhotoById(t *testing.T) {
	_, err := b.SendPhotoById(ID, "AgADAgADEaoxGx7ijwABUsJzu7MrSKBKDIMqAATYfJAUtDsU1ajtAQABAg", &SendPhotoOptions{Caption: "Caption for photo"})

	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

/*
func TestSendPhoto(t *testing.T) {
	_, err := b.SendPhotoById(ID, "toster.png", &SendPhotoOptions{Caption: "SendPhoto test"})

	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}
*/
