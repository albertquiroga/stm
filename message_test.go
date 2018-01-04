package main

import "testing"
import "strings"

func TestBuildURL(t *testing.T) {
  var builtURL = buildURL("hello", "there", "myfriend")
  var expectedURL = "https://api.telegram.org/bothello/sendMessage?chat_id=there&text=myfriend"

  if strings.Compare(builtURL, expectedURL) != 0 {
    t.Errorf("URL was incorrect, got: %s, want: %s.", builtURL, expectedURL)
  }
}
