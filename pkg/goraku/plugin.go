package goraku

import "github.com/nlopes/slack"

type Plugin interface {
	// Called when new message received
	Hear(bot Bot, message *slack.MessageEvent)
}
