package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"unicode"
)

var config interface{}

func main() {

	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("could not open config file")
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		panic("could not parse json from config file")
	}

	http.HandleFunc("/api/telegram", telegramHandler)
	http.HandleFunc("/api/slack", slackHandler)
	http.HandleFunc("/api/discord", discordHandler)
	http.HandleFunc("/api/wire", wireHandler)
	http.HandleFunc("/api/signal", signalHandler)
	http.HandleFunc("/api/matrix", matrixHandler)
	http.HandleFunc("/api/irc", ircHandler)
	_ = http.ListenAndServe(":8080", nil)
}

func telegramHandler(w http.ResponseWriter, r *http.Request) {
	outputMap := map[string]string{}
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

func slackHandler(w http.ResponseWriter, r *http.Request) {
	return
}
func discordHandler(w http.ResponseWriter, r *http.Request) {
	return
}
func wireHandler(w http.ResponseWriter, r *http.Request) {
	return
}
func signalHandler(w http.ResponseWriter, r *http.Request) {
	return
}
func matrixHandler(w http.ResponseWriter, r *http.Request) {
	return
}
func ircHandler(w http.ResponseWriter, r *http.Request) {
	return
}


func convertToMockery(str string) string {
	var convString = []rune{}
	for i := 0; i < len(str); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(2) != 0 {
			convString = append(convString, unicode.ToUpper(rune(str[i])))
		} else {
			convString = append(convString, unicode.ToLower(rune(str[i])))
		}
	}
	return string(convString)
}