# Telegram BOT

Implement golang interface for [Telegram Bot API](https://core.telegram.org/bots/api).

Implemented all methods up to: May 31, 2019 Bot API 4.3

# Examples

All examples are at: [example](example/) dir.

## Simple use

```go
package main

import (
	"context"
	"fmt"
	. "github.com/spyzhov/telego"
	"os"
	"strconv"
)

func main() {
	bot := New(os.Getenv("TOKEN"))
	chatId, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		panic(err)
	}
	user, err := bot.GetMe(context.Background())
	if err != nil {
		panic(err)
	}
	result, err := bot.SendMessage(context.Background(), &SendMessageRequest{
		ChatId: chatId,
		Text:   "Hello, my name is " + user.Username,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)
	if result.From != nil {
		fmt.Printf("From: %#v\n", result.From)
	}
	fmt.Printf("Chat: %#v\n", result.Chat)
}

```