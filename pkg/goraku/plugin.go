package goraku

import "github.com/nlopes/slack"

type Plugin interface {
	Hear(bot *Goraku, message slack.MessageEvent)
}
