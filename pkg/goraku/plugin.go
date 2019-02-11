package goraku

import "github.com/nlopes/slack"

type Plugin interface {
	Hear(bot Bot, message *slack.MessageEvent)
}
