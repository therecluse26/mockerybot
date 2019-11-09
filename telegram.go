package main

import (
	"context"
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"net/http"
	"github.com/aws/aws-lambda-go/lambda"
	"./mockery"
)

var config mockery.Config 

// Necessary to parse the message as sent to webhook
type msgContainer struct {
	Status  string           `json:"status"`
	Message tgbotapi.Message `json:"message"`
}

type lambdaEvent struct {
	Name string `json:"name"`
}

func main() {
	//config := mockery.GetConfigFromEnv()
	
	lambda.Start(telegramLambda)
}

func telegramLambda(ctx context.Context, name lambdaEvent) (string, error) {

	return "TESTING 123", nil

}


func telegramHandler(w http.ResponseWriter, r *http.Request) {

	var msg msgContainer
	data, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(data, &msg)

	if msg.Message.ReplyToMessage != nil && msg.Message.Command() == "mock" {

		bot, _ := tgbotapi.NewBotAPI(config["telegramApiKey"])

		replyMsg := mockery.ConvertToMockery(msg.Message.ReplyToMessage.Text)

		newMsg := tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
		bot.Send(newMsg)

	} else if msg.Message.IsCommand() {

		if msg.Message.Command() == "mock" {
			bot, _ := tgbotapi.NewBotAPI(config["telegramApiKey"])
			replyMsg := mockery.ConvertToMockery(msg.Message.CommandArguments())
			newMsg := tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)
		} else if msg.Message.Command() == "apologize" {
			bot, _ := tgbotapi.NewBotAPI(config["telegramApiKey"])
			replyMsg := mockery.ConvertToMockery("I'm sorry, " + msg.Message.CommandArguments() )
			newMsg := tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)
		}
	} 
}
