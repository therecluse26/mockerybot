package main

import (
	"encoding/json"
	"io/ioutil"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	"strings"
)

type telegramMessage struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int64    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int64    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int64    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int64    `json:"offset"`
			Length int64    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

func telegramHandler(w http.ResponseWriter, r *http.Request) {
	// If user is set, fetch the last message ID from them and mock it
		// TELEGRAM CURRENTLY DOESN'T SUPPORT THIS IN THEIR BOT API
	// If no user is set, mock the text string sent to mockerybot
	// If neither user nor message is set, mock the person who invoked mockerybot

	var mockery string	
	var msg telegramMessage
	data, err := ioutil.ReadAll(r.Body)
	//decoder := json.NewDecoder()
	err = json.Unmarshal(data, &msg)
	if err != nil {
		panic("Couldn't parse message")
	}

	bot, _ := tgbotapi.NewBotAPI(config["telegramApiKey"])


	/*if msg.Message.Entities != nil {
		fmt.Println(msg.Message.Entities[0].Type)

		chatConf := tgbotapi.ChatConfig{ChatID: msg.Message.Chat.ID, SuperGroupUsername: msg.Message.Chat.Type}
		chat, _ := bot.GetChat(chatConf)

		fmt.Println(chat)

		mockery = convertToMockery(msg.Message.Text)

	}*/ 

	cleanedText := strings.Replace(msg.Message.Text, "/mock", "", 1)
	cleanedText = strings.Replace(cleanedText, "@mockerybot", "", 1)

	mockery = convertToMockery(cleanedText)

	newMsg := tgbotapi.NewMessage(msg.Message.Chat.ID, mockery)

	
	//bot.Send(newMsg)
	//delMsg := tgbotapi.DeleteMessageConfig{msg.Message.Chat.ID, msg.Message.MessageID}
	//bot.DeleteMessage(delMsg)
	
}
