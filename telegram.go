package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
)

// Necessary to parse the message as sent to webhook
type msgContainer struct {
	Status  string           `json:"status"`
	Message tgbotapi.Message `json:"message"`
}

func telegramHandler(w http.ResponseWriter, r *http.Request) {

	var msg msgContainer
	data, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(data, &msg)

	if msg.Message.ReplyToMessage != nil && msg.Message.Command() == "mock" {

		bot, _ := tgbotapi.NewBotAPI(config["telegramApiKey"])

		replyMsg := convertToMockery(msg.Message.ReplyToMessage.Text)

		newMsg := tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
		bot.Send(newMsg)

	} else if msg.Message.IsCommand() {

		if msg.Message.Command() == "mock" {
			bot, _ := tgbotapi.NewBotAPI(config["telegramApiKey"])
			replyMsg := convertToMockery(msg.Message.CommandArguments())
			newMsg := tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)
	
		} else if msg.Message.Command() == "apologize" {
			bot, _ := tgbotapi.NewBotAPI(config["telegramApiKey"])
			replyMsg := convertToMockery("I'm sorry, " + msg.Message.CommandArguments() )
			newMsg := tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)
		}

		
		//delMsgCfg := tgbotapi.NewDeleteMessage(msg.Message.Chat.ID, msg.Message.MessageID)		
		//bot.DeleteMessage(delMsgCfg)
	} 
}
