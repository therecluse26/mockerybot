package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
)

func telegramHandler(w http.ResponseWriter, r *http.Request) {
	outputMap := map[string]string{}

	_, err := tgbotapi.NewBotAPI(config["telegramApiKey"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		outputMap["status"] = "error"
		outputMap["error"] = "Bot api connection broken"
		_ = json.NewEncoder(w).Encode(outputMap)
		return
	}






	user := r.URL.Query().Get("user")
	msgStr := r.URL.Query().Get("msg")
	if msgStr == "" && user == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		outputMap["status"] = "error"
		outputMap["error"] = "I need a user or a message to make fun of"
		_ = json.NewEncoder(w).Encode(outputMap)
		return
	}

	if user != "" {
		outputMap["user"] = user
	}
	if msgStr != "" {
		outputMap["msg"] = convertToMockery(msgStr)
	} else {
		outputMap["msg"] = " "//getUserMsg(user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputMap)




}