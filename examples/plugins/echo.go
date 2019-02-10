package plugins

import (
	"github.com/mtgto/goraku/pkg/goraku"
	"github.com/nlopes/slack"
)

type Echo struct {}

func (e *Echo) Hear(bot *goraku.Goraku, message slack.MessageEvent) {
	bot.Reply(message, message.Text)
}

var _ goraku.Plugin = (*Echo)(nil)
