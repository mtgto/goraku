package plugins

import (
	"github.com/mtgto/goraku/pkg/goraku"
	"github.com/nlopes/slack"
)

type Echo struct{}

func (e *Echo) Hear(bot goraku.Bot, message *slack.MessageEvent) {
	if message.User != bot.Me().ID {
		bot.Reply(message, message.Text)

		// Reply with using attachment
		//bot.ReplyWithOptions(message, slack.MsgOptionAsUser(true), slack.MsgOptionAttachments(slack.Attachment{
		//	Text:     message.Text,
		//	Fallback: message.Text,
		//}))
	}
}

var _ goraku.Plugin = (*Echo)(nil)
