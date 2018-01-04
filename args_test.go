package main

import "testing"
import "strings"

var expectedFileBotApiKey = "424848481:AAE0E8-1QX8Wf8LBhoChYLbLWiJjNvh_OqB"
var expectedFileChatId = "157875440"

func TestExtractArgsFromFile(t *testing.T) {
  var identityFilePath = "credentials.csv"
  botApiKey, chatId := extractArgsFromFile(identityFilePath)

  assertStringEquals(botApiKey, expectedFileBotApiKey, t)
  assertStringEquals(chatId, expectedFileChatId, t)
}

func TestParseFlags(t *testing.T) {
  // First case (no flags, a file, file should win)
  botApiKey, chatId := parseFlags("", "", "credentials.csv")
  assertStringEquals(botApiKey, expectedFileBotApiKey, t)
  assertStringEquals(chatId, expectedFileChatId, t)

  // Second case (one flag, a file, flag should win in its variable)
  botApiKey, chatId = parseFlags("", "12", "credentials.csv")
  assertStringEquals(botApiKey, expectedFileBotApiKey, t)
  assertStringEquals(chatId, "12", t)

  // Third case (both flags, a file, flags should win)
  botApiKey, chatId = parseFlags("34", "12", "credentials.csv")
  assertStringEquals(botApiKey, "34", t)
  assertStringEquals(chatId, "12", t)
}

func assertStringEquals(actual string, expected string, t *testing.T) {
  if strings.Compare(actual, expected) != 0 {
    t.Errorf("String did not match expected value, got: %s, want: %s.", actual, expected)
  }
}
