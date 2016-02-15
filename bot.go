package telegarobot

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram.
type Bot struct {
	Token string // Each bot is given a unique authentication token when it is created.
}

type response struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

func (bot *Bot) execApi(method string, values string) (respb []byte, err error) {
	resp, err := http.Get(fmt.Sprintf("%s%s/%s?%s", endpoint, bot.Token, method, values))
	resp.Close = true
	defer resp.Body.Close()
	respb, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
