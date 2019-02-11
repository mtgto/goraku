package plugins

import (
	"github.com/mtgto/goraku/pkg/goraku"
	"github.com/nlopes/slack"
	"testing"
)

type fakeGoraku struct {
	goraku.Bot
}

var receivedText string

func (g *fakeGoraku) Reply(message *slack.MessageEvent, text string) {
	receivedText = text
}

func TestPlugin_CheckMessage(t *testing.T) {
	echo := Echo{}
	bot := fakeGoraku{}
	text := "Sample Text"
	var testEvent = slack.MessageEvent{Msg: slack.Msg{Text: text}}
	echo.Hear(&bot, &testEvent)
	if receivedText != text {
		t.Fatal("No echo")
	}
}
