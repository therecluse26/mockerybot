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
			bot, _ := tgbotapi.NewBotAPI(config["apiKey"])
			replyMsg := mockery.ConvertToMockery(msg.Message.ReplyToMessage.Text)
			newMsg = tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)

		} else if msg.Message.Command() == "apologize" {
			bot, _ := tgbotapi.NewBotAPI(config["apiKey"])
			replyMsg := mockery.ConvertToMockery( mockery.MakeApology(msg.Message.ReplyToMessage.From.UserName) )
			newMsg = tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)
		}
		
	} else {

		if msg.Message.Command() == "mock" {
			bot, _ := tgbotapi.NewBotAPI(config["apiKey"])
			replyMsg := mockery.ConvertToMockery(msg.Message.CommandArguments())
			newMsg = tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)
			
		} else if msg.Message.Command() == "apologize" {
			bot, _ := tgbotapi.NewBotAPI(config["apiKey"])
			replyMsg := mockery.ConvertToMockery( mockery.MakeApology(msg.Message.CommandArguments()) )
			newMsg = tgbotapi.NewMessage(msg.Message.Chat.ID, replyMsg)
			bot.Send(newMsg)
		}
	} 

	return events.APIGatewayProxyResponse{Body: newMsg.Text, StatusCode: 200}, nil

}
