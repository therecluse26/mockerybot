package main
/*
import (
	"context"
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)



func initListeners(ctx context.Context, name lambdaEvent) (string, error) {
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
	http.HandleFunc("/api/keybase", keybaseHandler)
	http.HandleFunc("/api/slack", slackHandler)
	http.HandleFunc("/api/discord", discordHandler)
	http.HandleFunc("/api/wire", wireHandler)
	http.HandleFunc("/api/signal", signalHandler)
	http.HandleFunc("/api/matrix", matrixHandler)

	err = http.ListenAndServe(":"+config["port"], nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port " + string(config["port"]))

}

*/