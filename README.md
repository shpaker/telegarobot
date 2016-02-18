[![godoc badge](http://godoc.org/github.com/shpaker/telegarobot?status.png)](http://godoc.org/github.com/shpaker/telegarobot)
[![Build Status](https://drone.io/github.com/shpaker/telegarobot/status.png)](https://drone.io/github.com/shpaker/telegarobot/latest)
[![Coverage Status](https://coveralls.io/repos/github/shpaker/telegarobot/badge.svg?branch=master)](https://coveralls.io/github/shpaker/telegarobot?branch=master)
# telegarobot
 Go bindings for the Telegram Bot API 
*now in development*
# GET IT
```bash
go get github.com/shpaker/telegarobot
```
# USAGE
```go
package main

import (
	bot "github.com/shpaker/telegarobot"
	"log"
)

const (
	BOT_TOKEN          string = ""
	CHANNEL_OR_CHAT_ID string = ""
)

func main() {
	b := &bot.Bot{BOT_TOKEN}

	_, err := b.SendMessage(CHANNEL_OR_CHAT_ID, "_test_ *sendMessage* [link](google.com) `inline` ```outlone```", &bot.SendMessageOptions{ParseMode: bot.ParseModeMarkdown})
	if err != nil {
		log.Println(err)
	}
}
```