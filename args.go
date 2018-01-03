package main

import (
  "flag"
  "os"
  "fmt"
  "strings"
  "io/ioutil"
)

const botFlagName string = "bot"
const chatFlagName string = "cid"
const credentialsFileFlagName string = "i"

const defaultStringValue string = "replaceMePls"

// Gets the required botApiKey, chatId and text message to send from the CLI args.
func getParameters() (string, string, string) {
  checkForInsufficientArgs()
  botApiKey, chatId := parseFlags()
  textMessage := flag.Args()[0]
  checkForInvalidArgs(botApiKey, chatId)
  return botApiKey, chatId, textMessage
}

// Parses CLI flags and gives them a default value.
func parseFlags() (string, string) {
  var botApiKey string = defaultStringValue
  var chatId string = defaultStringValue

	var botApiKeyPtr = flag.String(botFlagName, defaultStringValue, "Telegram bot API key")
	var chatIdPtr = flag.String(chatFlagName, defaultStringValue, "Telegram chat identifier")
  var credentialsFilePtr = flag.String(credentialsFileFlagName, defaultStringValue, "Credentials (bot API key and chat ID) CSV file path")
  flag.Parse()

  if strings.Compare(*credentialsFilePtr, defaultStringValue) != 0 {
    botApiKey, chatId = extractArgsFromFile(*credentialsFilePtr)
  } else {
    botApiKey = *botApiKeyPtr
    chatId = *chatIdPtr
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
    panic(e)
  }
}

// Checks the number of CLI arguments. If there's only one (so no arguments since 1 is the program itself), prints
// the help message and stops execution.
func checkForInsufficientArgs() {
  if len(os.Args) == 1 {
    exitProgramWithHelp()
  }
}

// Ensures both the botApiKey and chatId have been set to a value different than the default one. If not, exits
// the program and shows the usage help message.
func checkForInvalidArgs(botApiKey string, chatId string) {
  if strings.Compare(botApiKey, defaultStringValue) == 0 || strings.Compare(chatId, defaultStringValue) == 0 {
    exitProgramWithHelp()
  }
}

// Prints the usage help message and halts execution
func exitProgramWithHelp() {
  printHelp()
  os.Exit(0)
}

// Prints a usage message for the user.
func printHelp() {
  fmt.Println("Usage: stm [options] message")
  fmt.Println("  options: check the man page.")
  fmt.Println("  message: the text message to be sent over Telegram. Remember to surround it in double quotes if it contains spaces.")
  fmt.Println("\nExample: stm -" + botFlagName + "=82AB71 -" + chatFlagName + "=193712 \"Hello, world!\"")
  fmt.Println("Example: stm -" + credentialsFileFlagName + "=credentials.csv \"Hello, world!\"")
}
