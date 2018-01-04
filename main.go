package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app              = kingpin.New("stm", "A command-line telegram message sender.")
	botApiKey        = app.Flag("bot", "Telegram bot API key").Short('b').String()
	chatId           = app.Flag("chat", "Telegram chat ID").Short('c').String()
	identityFilePath = app.Flag("file", "Path to identity CSV file containing the bot API key and the chat ID").Short('i').Default(os.Getenv("HOME") + "/.stm").String()
	message          = app.Arg("msg", "Message to send").Required().String()
)

func main() {
	app.Parse(os.Args[1:])
	botApiKey, chatId := getParameters(*botApiKey, *chatId, *identityFilePath)
	sendMessage(botApiKey, chatId, *message)
}
