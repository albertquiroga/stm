package main

import (
  "os"
  "fmt"
  "strings"
  "io/ioutil"
)

// Gets the required botApiKey, chatId and text message to send from the CLI args.
func getParameters(botApiKey string, chatId string, identityFilePath string) (string, string) {
  return parseFlags(botApiKey, chatId, identityFilePath)
}

func parseFlags(botApiKeyFlag string, chatIdFlag string, identityFilePathFlag string) (string, string) {
  botApiKey, chatId := extractArgsFromFile(identityFilePathFlag)
  if botApiKeyFlag != "" {
    botApiKey = botApiKeyFlag
  }
  if chatIdFlag != "" {
    chatId = chatIdFlag
  }
  return botApiKey, chatId
}

// Reads the file in the supplied filePath and extracts the botApiKey and chatId from it.
func extractArgsFromFile(filePath string) (string, string) {
  dat, err := ioutil.ReadFile(filePath)
  checkFileError(err)
  str := strings.Trim(string(dat), "\n")
  splitStr := strings.Split(str, ",")
  return splitStr[0], splitStr[1]
}

// Halts execution and prints stack trace if the file could not be read.
func checkFileError(e error) {
  if e != nil {
    fmt.Println(e)
    os.Exit(1)
  }
}
