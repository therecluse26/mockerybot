package main

import (
	"context"
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/therecluse26/mockerybot/mockery"
)

var config mockery.Config 

func main() {
	config = mockery.GetConfigFromEnv()
	lambda.Start(telegramLambda)
}

func telegramLambda(ctx context.Context, event map[string]interface{}) (events.APIGatewayProxyResponse, error) {

	var msg tgbotapi.Update 
	var newMsg tgbotapi.MessageConfig
	
	err := json.Unmarshal([]byte(event["body"].(string)), &msg)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Error: failed to parse JSON from request", StatusCode: 500}, err
	}
		
	if msg.Message.ReplyToMessage != (tgbotapi.Message{}.ReplyToMessage) {

		if msg.Message.Command() == "mock" {
			newMsg = botReply(msg.Message.Chat.ID, msg.Message.ReplyToMessage.Text)

		} else if msg.Message.Command() == "apologize" {
			newMsg = botReply(msg.Message.Chat.ID, mockery.MakeApology(msg.Message.ReplyToMessage.From.UserName))
		}
		
	} else {

		if msg.Message.Command() == "mock" {
			newMsg = botReply(msg.Message.Chat.ID, msg.Message.CommandArguments())

		} else if msg.Message.Command() == "apologize" {
			newMsg = botReply(msg.Message.Chat.ID, mockery.MakeApology(msg.Message.CommandArguments()))
		}
	} 

	// Necessary to clear out pending updates queue in webhook
	return events.APIGatewayProxyResponse{Body: newMsg.Text, StatusCode: 200}, nil
}

// Actually sends response to Telegram API
func botReply(chatId int64, sourceMsg string) tgbotapi.MessageConfig {
	bot, _ := tgbotapi.NewBotAPI(config["apiKey"])
	replyMsg := mockery.ConvertToMockery(sourceMsg)
	newMsg := tgbotapi.NewMessage(chatId, replyMsg)
	bot.Send(newMsg)
	return newMsg
}