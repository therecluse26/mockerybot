SHELL = /bin/bash

.PHONY: telegram
telegram:
	go build -o dist/telegram telegram.go
	zip dist/telegram.zip dist/telegram
	rm dist/telegram