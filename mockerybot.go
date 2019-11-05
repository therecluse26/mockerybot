package main

import (
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"unicode"
)

var config map[string]string

func main() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("could not open config file")
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		panic("could not parse json from config file")
	}
		
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "wElcOme to MOckErYboT")
	})
	http.HandleFunc("/api/telegram", telegramHandler)
	http.HandleFunc("/api/slack", slackHandler)
	http.HandleFunc("/api/discord", discordHandler)
	http.HandleFunc("/api/wire", wireHandler)
	http.HandleFunc("/api/signal", signalHandler)
	http.HandleFunc("/api/matrix", matrixHandler)

	fmt.Println("Listening on port " + string(config["port"]))
	_ = http.ListenAndServe(":"+config["port"], nil)
}

func convertToMockery(str string) string {
	var convertedStr []rune
	for i := 0; i < len(str); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(2) != 0 {
			convertedStr = append(convertedStr, unicode.ToUpper(rune(str[i])))
		} else {
			convertedStr = append(convertedStr, unicode.ToLower(rune(str[i])))
		}
	}
	return string(convertedStr)
}

