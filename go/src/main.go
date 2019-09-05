package main

import (
	"bot"
	"fmt"
)

func main() {
	msg := bot.LarkChatMsg{
		OpenId:  "ou_5ad573a6411d72b8305fda3a9c15c70e",
		RootId:  "om_40eb06e7b84dc71c03e009ad3c754195",
		UserId:  "92e39a99",
		Email:   "ll@abc.com",
		MsgType: "",
		Content: &bot.Content{
			Text: "hi",
		},
	}
	rmsg, _ := bot.Chat(msg)
	fmt.Println("#+v", rmsg)
}
