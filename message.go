package main

import (
  "net/http"
  "log"
)

const firstUrlPart string = "https://api.telegram.org/bot"
const secondUrlPart string = "/sendMessage?chat_id="
const finalUrlPart string = "&text="

// Sends the message through Telegram using the specified bot (botApiKey) to the specified user (chatId).
func sendMessage(botApiKey string, chatId string, textMessage string) {
  var url = buildURL(botApiKey, chatId, textMessage)
  sendGetRequest(url)
}

// Builds the URL to which send the GET request.
func buildURL(botApiKey string, chatId string, textMessage string) string {
	return firstUrlPart + botApiKey + secondUrlPart + chatId + finalUrlPart + textMessage
}

// Makes a GET request to the telegram endpoint, sending the message through the REST API.
func sendGetRequest(url string) {
  resp, err := http.Get(url)
  if err != nil {
  	log.Fatal(err)
  }
  defer resp.Body.Close()
}
