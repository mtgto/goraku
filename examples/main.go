package main

import (
	"flag"
	"github.com/mtgto/goraku/examples/plugins"
	"github.com/mtgto/goraku/pkg/goraku"
	"os"
)

func main() {
	var token string
	flag.StringVar(&token, "token", os.Getenv("SLACK_BOT_TOKEN"), "SlackのBotToken")
	flag.Parse()

	bot := goraku.NewSlackBot(token)
	bot.AddPlugin(&plugins.Echo{})

	// blocked until bot is disconnected.
	bot.Start()
}
