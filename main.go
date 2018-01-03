package main

func main() {
  botApiKey, chatId, textMessage := getParameters()
  sendMessage(botApiKey, chatId, textMessage)
}
